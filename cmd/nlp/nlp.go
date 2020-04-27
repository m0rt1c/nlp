package main

import (
	"bufio"
	"encoding/base64"
	"flag"
	"fmt"
	"log"
	"net/url"
	"os"
	"strconv"
	"strings"

	"github.com/chromium-netlog-parser/pkg/nlparser"
)

const (
	help = "help"
)

var (
	file    = flag.String("f", "", "File to parse")
	version = flag.Bool("version", false, "prints version")
	pc      = 0
)

func parseFile(s string) (nlparser.NetLog, error) {
	info, err := os.Stat(s)
	if err != nil {
		return nlparser.NetLog{}, err
	}
	if info.IsDir() {
		return nlparser.NetLog{}, fmt.Errorf("%s is not a valid file", s)
	}
	return nlparser.ParseNetLog(s)
}

func printhelp() {
	fmt.Println(`Commands:
    (help|?): show help text
    (parse|p) <file>: parse a new file
    (show|s) <arg>: show something of the parsed netlog (use show help for more)
    (extract|e) <arg>: extract events of the parsed netlog (use show help for more)
    (quit|q): exit`)
}

func handleShow(args []string, netlog *nlparser.NetLog) {
	switch args[0] {
	case "range":
		fmt.Printf("IDs range from 1 to %d\n", len(netlog.Events))
	case "id":
		if len(args) < 2 {
			return
		}
		index, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Println(err)
			return
		}
		e, ok := netlog.Events[int(index)]
		if !ok {
			fmt.Printf("IDs range from 0 to %d\n", len(netlog.Events))
		} else {
			fmt.Println(e)
		}
	case "all":
		for k := range netlog.Events {
			fmt.Println(netlog.Events[k])
		}
	case "next":
		pc = pc + 1
		if pc > len(netlog.Events) {
			pc = 1
		}
		fmt.Println(netlog.Events[pc])
	case "prev":
		pc = pc - 1
		if pc <= 0 {
			pc = len(netlog.Events)
		}
		fmt.Println(netlog.Events[pc])
	case "set":
		if len(args) < 2 {
			return
		}
		index, err := strconv.ParseInt(args[1], 10, 64)
		if err != nil {
			log.Println(err)
			return
		}
		pc = int(index)
	case help:
		fallthrough
	default:
		fmt.Println(`Options:
      range: events ids range
      id <number>: print event with given id
      all: dumps all events may. this will output lots of data
      next: print the next event based on print counter
      prev: print the prev event based on print counter
      set <number>: set the print counter`)
	}
}

func handleExtract(args []string, netlog *nlparser.NetLog) {
	switch args[0] {
	case "dns":
		for _, x := range netlog.FindDNSQueries() {
			fmt.Println(x)
		}
	case "url":
		for _, x := range netlog.FindURLRequests() {
			fmt.Println(x)
		}
	case "red":
		for _, x := range netlog.FindRedirections() {
			fmt.Println(x)
		}
	case "con":
		for _, x := range netlog.FindOpenedSocket() {
			fmt.Println(x)
		}
	case "src":
		if len(args) < 2 {
			fmt.Println("Not enough arguments")
		} else {
			_, err := os.Stat(args[1])
			if err != nil {
				err := os.Mkdir(args[1], os.ModePerm)
				if err != nil {
					fmt.Println(err)
					return
				}
			}
			fileCount := 0
			res := netlog.FindSources()
			for _, x := range res {
				u, err := url.Parse(x.ResourceName)
				if err != nil {
					fmt.Println(err)
					continue
				}
				parts := strings.Split(u.Path, "/")
				name := parts[len(parts)-1]
				if name == "" {
					name = fmt.Sprintf("index-%d", fileCount)
				}
				_, err = os.Stat(fmt.Sprintf("%s/%s", strings.TrimSuffix(args[1], "/"), u.Hostname()))
				if err != nil {
					os.Mkdir(fmt.Sprintf("%s/%s", strings.TrimSuffix(args[1], "/"), u.Hostname()), os.ModePerm)
				}
				filePath := fmt.Sprintf("%s/%s/%s", strings.TrimSuffix(args[1], "/"), u.Hostname(), name)
				file, err := os.Create(filePath)
				if err != nil {
					fmt.Printf("File %s was not created: %v\n", filePath, err)
					continue
				}
				out := []byte{}
				for _, part := range x.Base64EncodedBytes {
					decoded, err := base64.StdEncoding.DecodeString(part)
					if err != nil {
						fmt.Printf("Skipped some bytes while writing %s because: %v\n", filePath, err)
						continue
					}
					out = append(out, decoded...)
				}
				file.Write(out)
				file.Close()
				fileCount = fileCount + 1
			}
			fmt.Printf("Wrote %d out of %d files\n", fileCount, len(res))
		}
	case help:
		fallthrough
	default:
		fmt.Println(`Options:
      dns: print all dns queries
      url: print all url requests
      red: print all redirections
      con: print all connections
      src <path>: save all sources in the given directory. You also need to have captured the netlog with the --net-log-capture-mode=Everything flag.`)
	}
}

func main() {
	flag.Parse()
	var netlog nlparser.NetLog
	var err error

	if *version {
		fmt.Printf("version: %s\n", build.Version())
		return
	}

	if *file != "" {
		netlog, err = parseFile(*file)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("File parsed. Found %d events.\n", len(netlog.Events))
		}
	}

	running := true
	reader := bufio.NewReader(os.Stdin)
	for running {
		fmt.Print("> ")
		bytes, _, err := reader.ReadLine()
		if err != nil {
			running = false
		} else {
			pieces := strings.Split(fmt.Sprintf("%s", bytes), " ")
			if len(pieces) < 1 {
				fmt.Println("invalid command")
				continue
			}
			switch pieces[0] {
			case "?":
				fallthrough
			case help:
				printhelp()
			case "p":
				fallthrough
			case "parse":
				if len(pieces) < 2 {
					fmt.Println("invalid command")
				} else {
					netlog, err = parseFile(pieces[1])
					if err != nil {
						log.Println(err)
					} else {
						fmt.Printf("File parsed. Found %d events.\n", len(netlog.Events))
					}
				}
			case "s":
				fallthrough
			case "show":
				if len(pieces) < 2 {
					handleShow([]string{help}, &netlog)
				} else {
					handleShow(pieces[1:], &netlog)
				}
			case "e":
				fallthrough
			case "extract":
				if len(pieces) < 2 {
					handleExtract([]string{help}, &netlog)
				} else {
					handleExtract(pieces[1:], &netlog)
				}
			case "q":
				fallthrough
			case "quit":
				running = false
			default:
				fmt.Println("invalid command")
			}
		}
	}
}
