package grpc_ipc

import (
	"io"
	"log"
	"net"
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
	}(args)
	return err
}
