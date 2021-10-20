package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/brunocalza/go-libp2p-tutorial/protocols/echo"
	"github.com/libp2p/go-libp2p"
	peerlib "github.com/libp2p/go-libp2p-core/peer"
)

func main() {
	ctx := context.Background()

	port := os.Args[1]

	node, err := libp2p.New(ctx,
		libp2p.ListenAddrStrings("/ip4/127.0.0.1/tcp/"+port),
		libp2p.Ping(false),
		libp2p.NoSecurity,
	)
	if err != nil {
		panic(err)
	}

	fmt.Println("Listen addresses:", node.Addrs())

	peerInfo := peerlib.AddrInfo{
		ID:    node.ID(),
		Addrs: node.Addrs(),
	}
	addrs, err := peerlib.AddrInfoToP2pAddrs(&peerInfo)
	if err != nil {
		panic(err)
	}
	fmt.Println("libp2p node address:", addrs[0])

	protocol := echo.EchoProtocol{Host: node}
	node.SetStreamHandler(echo.ID, protocol.Handler)

	// wait for a SIGINT or SIGTERM signal
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, syscall.SIGINT, syscall.SIGTERM)
	<-ch
	fmt.Println("Received signal, shutting down...")
	time.Sleep(5 * time.Second)
	if err := node.Close(); err != nil {
		panic(err)
	}
}
