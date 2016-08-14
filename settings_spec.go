package grpc_ipc

// SettingsSpec is the top structure for mapping the JSON settings file.
type SettingsSpec struct {
	ListenAddress string   `json:"listen_address"`
	Pool          PoolSpec `json:"pool"`
}

// PoolSpec holds the information needed for initializing a Pool.
type PoolSpec struct {
	Size      int      `json:"pool_size"`
	Cmd       []string `json:"pool_cmd"`
	PortRange []int    `json:"pool_port_range"`
}
