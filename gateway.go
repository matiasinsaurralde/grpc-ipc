package main

import (
	"io"
	"log"
	"net"
)

// Gateway holds the basic gateway structure, the grpc.Server is contained here. It should be possible to create more than one "gateway" instance.
type Gateway struct {
	listener      *net.UnixListener
	listenAddress string
	pool          *Pool
}

// NewGateway initializes a new Gateway data structure.
func NewGateway(listenAddress string, initialPool *Pool) *Gateway {
	gw := &Gateway{
		listenAddress: listenAddress,
	}

	if initialPool != nil {
		gw.pool = initialPool
	} else {
		gw.pool = &Pool{}
	}

	return gw
}

// Serve starts the internal grpc server
func (g *Gateway) Serve() (err error) {
	g.pool.Start()

	var gatewayAddress *net.UnixAddr
	gatewayAddress, err = net.ResolveUnixAddr("unix", g.listenAddress)

	g.listener, err = net.ListenUnix("unix", gatewayAddress)
	defer g.listener.Close()

	for {
		conn, err := g.listener.AcceptUnix()
		defer conn.Close()

		if err != nil {
			log.Println("Error:", err)
		}

		go g.handle(conn)
	}

	return err
}

// Handle connection I/O
func (g *Gateway) handle(conn *net.UnixConn) {
	worker := g.pool.PickWorker()
	// log.Println("Picking worker", worker.addr)
	proxyConn, _ := net.DialUnix("unix", nil, worker.addr)
	go copyAndClose(proxyConn, conn)
	go copyAndClose(conn, proxyConn)
}

// As seen in github.com/elazarl/goproxy
func copyAndClose(dst, src *net.UnixConn) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Println("Error:", err)
	}
	dst.CloseWrite()
	src.CloseRead()
}
