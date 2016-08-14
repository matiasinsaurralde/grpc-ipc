package main

type SettingsSpec struct {
  PoolSize  int  `json:"pool_size"`
  Listen  string `json:"listen_port"`
}
