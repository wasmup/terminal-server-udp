package main

import (
	"log"
	"net"
	"os"
)

func main() {
	pc, err := net.ListenPacket("udp", ":8080")
	if err != nil {
		log.Fatal(err)
	}
	defer pc.Close()
	buf := make([]byte, 1024)
	for {
		n, _, err := pc.ReadFrom(buf)
		if err != nil {
			break
		}
		os.Stdout.Write(buf[:n])
	}
}
