package main

import(
  "log"
  "io/ioutil"
  "encoding/json"
)

var Settings *SettingsSpec

func init() {
  log.Println("Starting grpc-ipc")

  var settingsData []byte
  var err error

  settingsData, err = ioutil.ReadFile("settings.json")
  if err != nil {
    panic(err)
  }

  Settings = &SettingsSpec{}

  json.Unmarshal(settingsData, Settings)
}

func main() {

}
