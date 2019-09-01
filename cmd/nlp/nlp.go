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

func handleShow(args []string, netlog *nlparser.NetLog) {}

func main() {
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
					fmt.Println("invalid command")
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
