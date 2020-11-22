package main

import (
	"net"
)

func main() {
	addr := ":8888"
	protocol := "udp"
	packetConn, err := net.ListenPacket(protocol, addr)
	if err != nil {

	}
}
