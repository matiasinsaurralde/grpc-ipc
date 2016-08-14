package main

import(
  "net"
  "google.golang.org/grpc"
)

type Gateway struct {
  grpc *grpc.Server
  listener net.Listener
  listenAddress string
}

func NewGateway(listenAddress string) Gateway {
  gw := Gateway{
    grpc: grpc.NewServer(),
    listenAddress: listenAddress,
  }
  return gw
}

func(g *Gateway) Serve() (err error) {
  g.listener, err = net.Listen( "tcp", g.listenAddress )
  g.grpc = grpc.NewServer()
  return err
}
