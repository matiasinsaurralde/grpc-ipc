package main

import(
  "log"
  "io/ioutil"
  "encoding/json"
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
  gateway := NewGateway(settings.ListenAddress)
  err := gateway.Serve()
  if err != nil {
    panic(err)
  }
}
