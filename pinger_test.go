package main

import (
	"fmt"
	"net"
	"testing"
	"time"
)

func TestOneHost(t *testing.T) {
	expected := 1
	res, err := Hosts("127.0.0.1/32")
	if err != nil {
		t.Errorf("error: %v", err.Error())
	}
	if len(res) != expected {
		t.Errorf("should be %d host, but got %d", expected, len(res))
	}
}

func TestMultipleHosts(t *testing.T) {
	expected := 254
	res, err := Hosts("10.0.0.0/24")
	if err != nil {
		t.Errorf("error: %v", err.Error())
	}
	if len(res) != expected {
		t.Errorf("should be %d host, but got %d", expected, len(res))
	}
}

func TestHostAlive(t *testing.T) {
	hosts := 1
	ip := "127.0.0.1"
	pingChan := make(chan string, 1)
	pongChan := make(chan Pong, hosts)
	doneChan := make(chan []Pong)
	go ping(pingChan, pongChan)
	go receivePong(1, pongChan, doneChan)
	pingChan <- ip
	res := <-doneChan
	if len(res) != 1 {
		t.Errorf("should be %d hosts, but got %d", hosts, len(res))
	}
}

func TestOpenPort(t *testing.T) {
	ip := "127.0.0.1"
	hosts := 1
	port := 1111
	_, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		t.Errorf("failed to open port %d", port)
	}
	ScanPort(ip, port, 500*time.Millisecond)
	if len(openPorts[ip]) != hosts {
		t.Errorf("should 1 port, but got %d", len(openPorts[ip]))
	}

	if openPorts[ip][0] != port {
		t.Errorf("port should be %d, but got %d", port, openPorts[ip][0])
	}
}
