package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/chromium-netlog-parser/pkg/nlparser"
)

var (
	file = flag.String("file", "", "File to parsed")
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

func printHelp() {
	fmt.Println(`Commands:
    help, ?: show help text
    parse <file>: parse a new file
    show <arg>: show something of the parsed netlog (use show help for more)
    quit, q: exit`)
}

func handleShow(args []string, netlog *nlparser.NetLog) {
	switch args[0] {
	case "help":
		fallthrough
	default:
		fmt.Println(`Options:
      range: events ids range
      <number>: print event with given id
      all: dumps all events may. this will output lots of data
      next: print next event based on print counter
      prev: print prev event based on print counter
      set <number>: set print counter`)
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
			case "help":
				printHelp()
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
					handleShow([]string{"help"}, &netlog)
				} else {
					handleShow(pieces[1:], &netlog)
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
