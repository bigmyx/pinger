package main

import (
	"net"
)

func Hosts(cidr string) ([]string, error) {

	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	if len(ips) > 1 {
		return ips[1 : len(ips)-1], nil
	}

	return ips, nil

}

func inc(ip net.IP) {

	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}

}
