package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"golang.org/x/sync/semaphore"
)

const (
	startPort = 80
	endPort   = 10000
)

func Usage() {
	fmt.Printf("Usage: %s [options] <cidr>\noptions:\n", os.Args[0])
	flag.PrintDefaults()
}

func main() {

	start := time.Now()
	flag.Usage = Usage

	maxThreads := flag.Int("threads", 50, "# of threads")
	flag.Parse()
	if flag.NArg() != 1 {
		Usage()
		log.Fatal("missing args")
	}
	ipRange := flag.Arg(0)

	sp := strings.Split(ipRange, "/")

	if len(sp) < 2 {
		ipRange = ipRange + "/32"
	}

	hosts, err := Hosts(ipRange)
	if err != nil {
		log.Fatal(err.Error())
	}

	pingChan := make(chan string, *maxThreads)
	pongChan := make(chan Pong, len(hosts))
	doneChan := make(chan []Pong)

	for i := 0; i < *maxThreads; i++ {
		go ping(pingChan, pongChan)
	}

	go receivePong(len(hosts), pongChan, doneChan)

	for _, ip := range hosts {
		pingChan <- ip
	}

	alives := <-doneChan

	for _, tgt := range alives {

		ps := &PortScanner{
			ip:   tgt.IP,
			lock: semaphore.NewWeighted(Ulimit()),
		}
		ps.Start(startPort, endPort, 500*time.Millisecond)

		colorize(ColorBlue, fmt.Sprintf("Report for host %s", tgt.IP))

		for _, port := range openPorts[tgt.IP] {
			colorize(ColorRed, fmt.Sprintf("port %s\topen", strconv.Itoa(port)))
		}

	}
	colorize(ColorYellow,
		fmt.Sprintf("Scanned %d hosts in %s",
			len(alives), time.Since(start)))
}
