package config

type StreamClientConfig struct {
	// from client
	DestService               string                     `yaml:"dest_service"`
	HostPorts                 []string                   `yaml:"host_ports"`
	ConnectTimeout            *TimeInterval              `yaml:"connect_timeout"`
	Tag                       *Tag                       `yaml:"tag"`
	StatsLevel                int                        `yaml:"stats_level"`
	ConnReporterEnabled       bool                       `yaml:"conn_reporter_enabled"`
	GRPCConnPoolSize          uint32                     `yaml:"grpc_conn_pool_size"`
	GRPCWriteBufferSize       uint32                     `yaml:"grpc_write_buffer_size"`
	GRPCReadBufferSize        uint32                     `yaml:"grpc_read_buffer_size"`
	GRPCInitialWindowSize     uint32                     `yaml:"grpc_initial_window_size"`
	GRPCInitialConnWindowSize uint32                     `yaml:"grpc_initial_conn_window_size"`
	GRPCMaxHeaderListSize     uint32                     `yaml:"grpc_max_header_list_size"`
	GRPCKeepaliveParams       *GRPCClientKeepaliveParams `yaml:"grpc_keepalive_params"`
	// from client advanced
	ClientBasicInfo *EndpointBasicInfo `yaml:"client_basic_info"`
}
