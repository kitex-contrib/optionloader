package config

type CalloptConfig struct {
	HostPort       string        `yaml:"host_port"`
	URL            string        `yaml:"url"`
	HTTPHost       string        `yaml:"http_host"`
	RPCTimeout     *TimeInterval `yaml:"rpc_timeout"`
	ConnectTimeout *TimeInterval `yaml:"connect_timeout"`
	Tag            *Tag          `yaml:"tag"`
	GRPCCompressor string        `yaml:"grpc_compressor"`
}
