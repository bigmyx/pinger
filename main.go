package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jedib0t/go-pretty/table"
	"golang.org/x/sync/semaphore"
)

const (
	startPort = 80
	endPort   = 10000
)

func main() {
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("missing args")
	}
	ipRange := flag.Arg(0)

	hosts, err := Hosts(ipRange)
	if err != nil {
		log.Fatal(err.Error())
	}

	maxThreads := 100
	pingChan := make(chan string, maxThreads)
	pongChan := make(chan Pong, len(hosts))
	doneChan := make(chan []Pong)

	for i := 0; i < maxThreads; i++ {
		go ping(pingChan, pongChan)
	}

	go receivePong(len(hosts), pongChan, doneChan)

	for _, ip := range hosts {
		pingChan <- ip
	}

	alives := <-doneChan

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)

	for _, tgt := range alives {

		ps := &PortScanner{
			ip:   tgt.IP,
			lock: semaphore.NewWeighted(Ulimit()),
		}
		ps.Start(startPort, endPort, 500*time.Millisecond)

		t.AppendRow(table.Row{tgt.IP, ""})
		for _, port := range openPorts[tgt.IP] {
			t.AppendRow(table.Row{"", port})
		}
		t.AppendSeparator()

	}
	scanned := fmt.Sprintf("found %d hosts\n", len(alives))
	t.AppendFooter(table.Row{scanned})
	t.Render()
}
