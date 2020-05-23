package main

import (
	"flag"
	"log"
)

func main() {
	// Config read
	var cfg Config
	readFile(&cfg)
	readEnv(&cfg)

	// Parse command args
	flag.Parse()
	if flag.NArg() != 1 {
		log.Fatal("missing args")
	}
	ipRange := flag.Arg(0)

	hosts, err := Hosts(ipRange)
	if err != nil {
		log.Fatal(err.Error())
	}

	pingChan := make(chan string, cfg.Pinger.Threads)
	pongChan := make(chan Pong, len(hosts))
	doneChan := make(chan []Pong)

	for i := 0; i < cfg.Pinger.Threads; i++ {
		go ping(pingChan, pongChan)
	}

	go receivePong(len(hosts), pongChan, doneChan)

	for _, ip := range hosts {
		pingChan <- ip
	}

	alives := <-doneChan
  report(alives)
	//pp.Println(alives)

}
