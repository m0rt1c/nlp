package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/AndreaJegher/chromium-netlog-parser/pkg/build"
	"github.com/AndreaJegher/chromium-netlog-parser/pkg/nlparser"
)

const (
	help = "help"
)

var (
	file    = flag.String("f", "", "File to parse")
	version = flag.Bool("version", false, "Prints command version")
	command = flag.String("c", "", "Pass command as a string instead of reading from stdin")
	pc      = 0
)

func main() {
	flag.Parse()
	var netlog nlparser.NetLog
	var err error

	if *version {
		fmt.Printf("version: %s\n", build.Version())
		return
	}

	if *file != "" {
		netlog, err = nlparser.ParseFile(*file)
		if err != nil {
			log.Fatal(err)
		} else {
			fmt.Printf("File parsed. Found %d events.\n", len(netlog.Events))
		}
	}

	running := true
	reader := bufio.NewReader(os.Stdin)
	for running {
		bytes, _, err := reader.ReadLine()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Print("> ")
		out, err := nlparser.HandleCommand(fmt.Sprintf("%s", bytes), netlog)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Prinltn(out)
	}
}
