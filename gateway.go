package main

import (
	"google.golang.org/grpc"
	"net"
)

// Gateway holds the basic gateway structure, the grpc.Server is contained here. It should be possible to create more than one "gateway" instance.
type Gateway struct {
	grpc          *grpc.Server
	listener      net.Listener
	listenAddress string
  pool *Pool
}

// NewGateway initializes a new Gateway data structure.
func NewGateway(listenAddress string, initialPool *Pool) *Gateway {
	gw := &Gateway{
		grpc:          grpc.NewServer(),
		listenAddress: listenAddress,
	}

  if initialPool != nil {
    gw.pool =  initialPool
  } else {
    gw.pool = &Pool{}
  }

	return gw
}

// Serve starts the internal grpc server
func (g *Gateway) Serve() (err error) {
  g.pool.Start()
	g.listener, err = net.Listen("tcp", g.listenAddress)
	g.grpc = grpc.NewServer()
	return err
}
