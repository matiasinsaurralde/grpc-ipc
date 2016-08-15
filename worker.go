package grpc_ipc

import (
	"io"
	"log"
	"net"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/dchest/uniuri"
)

const uriPrefix string = "grpc_ipc_"

// Worker represents a single process in the pool.
type Worker struct {
	pool   *Pool
	addr   *net.UnixAddr
	conn *net.UnixConn
	output io.ReadCloser
}

// Prepare will set the socket path.
func (w *Worker) Prepare() (err error) {
	socketFile := strings.Join([]string{uriPrefix, uniuri.New()}, "")
	socketPath := path.Join("/tmp", socketFile)
	w.addr, err = net.ResolveUnixAddr("unix", socketPath)

	return err
}

// Start starts the worker process.
func (w *Worker) Start() (err error) {
	args := make([]string, len(w.pool.Cmd)+1)
	copy(args, w.pool.Cmd)
	args = append(args, w.addr.Name)

	go func(args []string) {
		cmd := exec.Command(args[0], args[1:]...)
		w.output, err = cmd.StdoutPipe()
		if err != nil {
			log.Println("Error:", err)
		}
		if err := cmd.Start(); err != nil {
			log.Println("Error:", err)
		}
		if err := cmd.Wait(); err != nil {
			log.Println("Error:", err)
		}
		return
	}(args)

	go func() {
		var err error
		for {
			_, err = os.Stat(w.addr.Name)
			if err == nil {
				break
			}
		}
		for {
			w.conn, err = net.DialUnix("unix", nil, w.addr)
			if err == nil && w.conn != nil {
				break
			}
		}

		log.Println("Worker connection established:", *w.conn)
		return
	}()

	return err
}
