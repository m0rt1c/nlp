package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/chromium-netlog-parser/pkg/nlparser"
)

const (
	help = "help"
)

var (
	file = flag.String("file", "", "File to parsed")
	pc   = 0
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
    help, ?: show help text
    parse <file>: parse a new file
    show <arg>: show something of the parsed netlog (use show help for more)
    quit, q: exit`)
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
      next: print next event based on print counter
      prev: print prev event based on print counter
      set <number>: set print counter`)
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
		for _, x := range netlog.FindSources() {
			fmt.Println(x)
		}
	case help:
		fallthrough
	default:
		fmt.Println(`Options:
      dns: print all dns queries
      url: print all url requests
      red: print all redirections
      con: print all connections
      src: print all sources. this may be a lot of data. you also need to have capture the netlog with --net-log-capture-mode=Everything flag.`)
	}
}

func main() {
	flag.Parse()
	var netlog nlparser.NetLog
	var err error

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
			case "show":
				if len(pieces) < 2 {
					handleShow([]string{help}, &netlog)
				} else {
					handleShow(pieces[1:], &netlog)
				}
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
