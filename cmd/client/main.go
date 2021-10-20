package main

import (
	"context"
	"os"
	"time"

	"github.com/brunocalza/go-libp2p-tutorial/protocols/echo"
	"github.com/libp2p/go-libp2p"
	peerlib "github.com/libp2p/go-libp2p-core/peer"
	"github.com/multiformats/go-multiaddr"
)

func main() {
	ctx := context.Background()

	peerAddress := os.Args[1]
	inputWord := os.Args[2]

	node, err := libp2p.New(ctx,
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/0"),
		libp2p.Ping(false),
		libp2p.NoSecurity,
	)
	if err != nil {
		panic(err)
	}

	addr, err := multiaddr.NewMultiaddr(peerAddress)
	if err != nil {
		panic(err)
	}
	peer, err := peerlib.AddrInfoFromP2pAddr(addr)
	if err != nil {
		panic(err)
	}
	if err := node.Connect(ctx, *peer); err != nil {
		panic(err)
	}

	protocol := echo.EchoProtocol{node}
	protocol.Echo(ctx, peer.ID, inputWord)

	time.Sleep(5 * time.Second)
	if err := node.Close(); err != nil {
		panic(err)
	}
}
