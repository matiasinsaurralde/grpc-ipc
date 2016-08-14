package main

import(
  "fmt"
  "net/http"
  "encoding/json"
  "os"
  "os/signal"
  "syscall"
  "net"
  "time"
  "log"

  ipc "github.com/matiasinsaurralde/grpc-ipc"

  "golang.org/x/net/context"
  "google.golang.org/grpc"
  pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

const (
	address     = "/tmp/gateway"
	defaultName = "world"
)

func dialer(addr string, timeout time.Duration) (net.Conn, error) {
	return net.DialTimeout("unix", addr, timeout)
}

func handleHello(w http.ResponseWriter, req *http.Request) {
  conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithDialer(dialer))
  if err != nil {
    log.Fatalf("did not connect: %v", err)
  }
  defer conn.Close()
  c := pb.NewGreeterClient(conn)

  // Contact the server and print out its response.
  name := defaultName
  if len(os.Args) > 1 {
    name = os.Args[1]
  }
  r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
  if err != nil {
    log.Fatalf("could not greet: %v", err)
  }

  jsonOutput, _ := json.Marshal(r)

  w.Write(jsonOutput)
}

func main() {
  fmt.Println("main()")

  settings, _ := ipc.LoadConfig("settings.json")

  pool := ipc.NewPool(settings.Pool.Size, settings.Pool.Cmd, settings.Pool.PortRange)
  gateway := ipc.NewGateway(settings.ListenAddress, pool)

  go gateway.Serve()

  http.HandleFunc("/", handleHello)
  go http.ListenAndServe(":8000", nil)

  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
  <-c

  os.Remove("/tmp/gateway")
}
