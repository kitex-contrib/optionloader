package config

type ClientConfig struct {
	// basic
	TransportProtocol         int                        `yaml:"transport_protocol"`
	DestService               string                     `yaml:"dest_service"`
	HostPorts                 []string                   `yaml:"host_ports"`
	LongConnection            *IdleConfig                `yaml:"long_connection"`
	MuxConnection             int                        `yaml:"mux_connection"`
	RPCTimeout                *TimeInterval              `yaml:"rpc_timeout"`
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
	// advanced
	HTTPConnection  bool               `yaml:"http_connection"`
	ClientBasicInfo *EndpointBasicInfo `yaml:"client_basic_info"`
}

type IdleConfig struct {
	MinIdlePerAddress int          `yaml:"min_idle_per_address"`
	MaxIdlePerAddress int          `yaml:"max_idle_per_address"`
	MaxIdleGlobal     int          `yaml:"max_idle_global"`
	MaxIdleTimeout    TimeInterval `yaml:"max_idle_timeout"`
}

type GRPCClientKeepaliveParams struct {
	Time                TimeInterval `yaml:"time"`
	Timeout             TimeInterval `yaml:"timeout"`
	PermitWithoutStream bool         `yaml:"permit_without_stream"`
}
