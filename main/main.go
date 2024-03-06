package main

import (
	"Distributed-CAS-Storage/p2p"
	"fmt"
	"log"
)

func main() {
	tr := p2p.NewTCPTransport(":3000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello")
}
