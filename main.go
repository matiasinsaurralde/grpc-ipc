package main

import (
	"encoding/json"
	"io/ioutil"
  "os"
  "os/signal"
  "syscall"
	"log"
)

var settings *SettingsSpec

func init() {
	log.Println("Starting grpc-ipc")

	var settingsData []byte
	var err error

	settingsData, err = ioutil.ReadFile("settings.json")
	if err != nil {
		panic(err)
	}

	settings = &SettingsSpec{}

	json.Unmarshal(settingsData, settings)
}

func main() {
  pool := NewPool( settings.Pool.Size, settings.Pool.Cmd, settings.Pool.PortRange )
	gateway := NewGateway( settings.ListenAddress, pool )
  
	go gateway.Serve()

  c := make(chan os.Signal, 1)
  signal.Notify(c, os.Interrupt, os.Kill, syscall.SIGTERM)
  <-c

  os.Remove("/tmp/gateway")
}
