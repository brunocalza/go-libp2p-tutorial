package echo

import (
	"bufio"
	"context"
	"io/ioutil"
	"log"

	"github.com/libp2p/go-libp2p-core/host"
	"github.com/libp2p/go-libp2p-core/network"
	"github.com/libp2p/go-libp2p-core/peer"
)

const ID = "echo/1.0.0"

// EchoProtocol implements
type EchoProtocol struct {
	Host host.Host
}

// Handler handles an incoming stream of bytes of echoing it back
func (p *EchoProtocol) Handler(s network.Stream) {
	if err := doEcho(s); err != nil {
		log.Println(err)
		s.Reset()
	} else {
		s.Close()
	}
}

// Echo receives a word as input and sends to a peer and logs the response back
func (p *EchoProtocol) Echo(ctx context.Context, pid peer.ID, word string) {
	s, err := p.Host.NewStream(ctx, pid, ID)
	if err != nil {
		log.Println(err)
		return
	}
	_, err = s.Write([]byte(word + "\n"))
	if err != nil {
		log.Println(err)
		return
	}

	out, err := ioutil.ReadAll(s)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("read reply: %q\n", out)
}

// doEcho reads a line of data a stream and writes it back
func doEcho(s network.Stream) error {
	buf := bufio.NewReader(s)
	str, err := buf.ReadString('\n')
	if err != nil {
		return err
	}

	log.Printf("read: %s", str)
	_, err = s.Write([]byte(str))
	return err
}
