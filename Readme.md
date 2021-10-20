# libp2p go example

## Description

Toy example to get a sense of how [libp2p](https://libp2p.io/) works. I've separated the example into pieces to make it clear how each part works. There are two binaries: one that works as a server and another as a client. Both communicate using the `echo` protocol.

Mostly inspired by [Getting started with go-libp2p](https://docs.libp2p.io/tutorials/getting-started/go/) and [libp2p go example](https://github.com/libp2p/go-libp2p/blob/master/examples/echo/main.go).

## How to build

`go build ./cmd/server/`

`go build ./cmd/client/`

## How to run

`./server [PORT]`

`./client [PEER] [INPUT]`
