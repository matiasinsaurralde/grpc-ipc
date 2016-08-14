package grpc_ipc

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// var settings *SettingsSpec

func init() {
	log.Println("Starting grpc-ipc")
}

func LoadConfig(path string) (*SettingsSpec, error) {
  var settingsData []byte
  var err error

  settingsData, err = ioutil.ReadFile(path)
  if err != nil {
    panic(err)
  }

  var settings *SettingsSpec
  err = json.Unmarshal(settingsData, &settings)

  return settings, err
}

func main() {
  var settings *SettingsSpec

  settings, _ = LoadConfig("settings.json")

	pool := NewPool(settings.Pool.Size, settings.Pool.Cmd, settings.Pool.PortRange)
	gateway := NewGateway(settings.ListenAddress, pool)

	go gateway.Serve()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
	<-c

	os.Remove("/tmp/gateway")
}
