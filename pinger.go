package main

import (
	"os/exec"
)

type Pong struct {
	IP    string
	Alive bool
}

func ping(pingChan <-chan string, pongChan chan<- Pong) {

	for ip := range pingChan {
		_, err := exec.Command("ping", "-c1", "-t10", ip).Output()
		var alive bool
		if err != nil {
			alive = false
		} else {
			alive = true
		}
		pongChan <- Pong{IP: ip, Alive: alive}
	}

}

func receivePong(pongNum int, pongChan <-chan Pong, doneChan chan<- []Pong) {

	var alives []Pong
	for i := 0; i < pongNum; i++ {
		pong := <-pongChan
		if pong.Alive {
			alives = append(alives, pong)
		}
	}
	doneChan <- alives

}
