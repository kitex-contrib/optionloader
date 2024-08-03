package config

type ServerConfig struct {
	// basic
	MuxTransport                    bool                                  `yaml:"mux_transport"`                       // For WithMuxTransport
	ReadWriteTimeout                *TimeInterval                         `yaml:"read_write_timeout"`                  // For WithReadWriteTimeout
	ExitWaitTime                    *TimeInterval                         `yaml:"exit_wait_time"`                      // For WithExitWaitTime
	MaxConnIdleTime                 *TimeInterval                         `yaml:"max_conn_idle_time"`                  // For WithMaxConnIdleTime
	StatsLevel                      int                                   `yaml:"stats_level"`                         // For WithStatsLevel
	ServiceAddr                     *NetAddr                              `yaml:"service_addr"`                        // For WithServiceAddr
	GRPCWriteBufferSize             uint32                                `yaml:"grpc_write_buffer_size"`              // For WithGRPCWriteBufferSize
	GRPCReadBufferSize              uint32                                `yaml:"grpc_read_buffer_size"`               // For WithGRPCReadBufferSize
	GRPCInitialWindowSize           uint32                                `yaml:"grpc_initial_window_size"`            // For WithGRPCInitialWindowSize
	GRPCInitialConnWindowSize       uint32                                `yaml:"grpc_initial_conn_window_size"`       // For WithGRPCInitialConnWindowSize
	GRPCKeepaliveParams             *GRPCServerKeepaliveParams            `yaml:"grpc_keepalive_params"`               // For WithGRPCKeepaliveParams
	GRPCKeepaliveEnforcementPolicy  *GRPCServerKeepaliveEnforcementPolicy `yaml:"grpc_keepalive_enforcement_policy"`   // For WithGRPCKeepaliveEnforcementPolicy
	GRPCMaxConcurrentStreams        uint32                                `yaml:"grpc_max_concurrent_streams"`         // For WithGRPCMaxConcurrentStreams
	GRPCMaxHeaderListSize           uint32                                `yaml:"grpc_max_header_list_size"`           // For WithGRPCMaxHeaderListSize
	ContextBackup                   *ContextBackup                        `yaml:"context_backup"`                      // For WithContextBackup
	RefuseTrafficWithoutServiceName bool                                  `yaml:"refuse_traffic_without_service_name"` // For WithRefuseTrafficWithoutServiceName
	EnableContextTimeout            bool                                  `yaml:"enable_context_timeout"`              // For WithEnableContextTimeout
	// advanced
	ServerBasicInfo *EndpointBasicInfo `yaml:"server_basic_info"` // For WithServerBasicInfo
	ReusePort       bool               `yaml:"reuse_port"`        // For WithReusePort
	// stream
	CompatibleMiddlewareForUnary bool `yaml:"compatible_middleware_for_unary"` // For WithCompatibleMiddlewareForUnary
}

// NetAddr defines the network and address of the service.
// Network name of the network (for example, "tcp", "udp").
// Address string form of address (for example, "192.0.2.1:25", "[2001:db8::1]:80").
type NetAddr struct {
	Network string `yaml:"network"`
	Address string `yaml:"address"`
}

type GRPCServerKeepaliveParams struct {
	MaxConnectionIdle     TimeInterval `yaml:"max_connection_idle"`
	MaxConnectionAge      TimeInterval `yaml:"max_connection_age"`
	MaxConnectionAgeGrace TimeInterval `yaml:"max_connection_age_grace"`
	Time                  TimeInterval `yaml:"time"`
	Timeout               TimeInterval `yaml:"timeout"`
}

type GRPCServerKeepaliveEnforcementPolicy struct {
	MinTime             TimeInterval `yaml:"min_time"`
	PermitWithoutStream bool         `yaml:"permit_without_stream"`
}

type ContextBackup struct {
	Enable bool `yaml:"enable"`
	Async  bool `yaml:"async"`
}
