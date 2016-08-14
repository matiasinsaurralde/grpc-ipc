package main

type SettingsSpec struct {
	ListenAddress string   `json:"listen_address"`
	Pool          PoolSpec `json:"pool"`
}

type PoolSpec struct {
	PoolSize      int      `json:"pool_size"`
	PoolCmd       []string `json:"pool_cmd"`
	PoolPortRange []int    `json:"pool_port_range"`
}
