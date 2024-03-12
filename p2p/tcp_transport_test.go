package p2p

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTCPTransport(t *testing.T) {
	//listenAddr := ":4000"
	opts := TCPTransportOpts{
		ListenAddr:    "3000",
		HandShakeFunc: NOPHandShakeFunc,
		Decoder:       DefaultDecoder{},
	}
	tr := NewTCPTransport(opts)
	assert.Equal(t, tr.ListenAddr, ":3000")
	assert.Nil(t, tr.ListenAndAccept())
	select {}
}
