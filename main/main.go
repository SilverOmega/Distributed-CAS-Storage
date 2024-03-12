package main

import (
	"Distributed-CAS-Storage/p2p"
	"fmt"
	"log"
)

func OnPeer(peer p2p.Peer) error {
	peer.Close()
	return nil
}
func main() {
	tcpOpts := p2p.TCPTransportOpts{
		ListenAddr:    "3000",
		HandShakeFunc: p2p.NOPHandShakeFunc,
		Decoder:       p2p.DefaultDecoder{},
		OnPeer:        OnPeer,
	}
	tr := p2p.NewTCPTransport(tcpOpts)
	go func() {
		for {
			msg := <-tr.Consume()
			fmt.Printf("%+v/n", msg)
		}
	}()
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
}
