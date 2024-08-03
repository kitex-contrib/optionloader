package config

type StreamCallConfig struct {
	// from callopt
	HostPort       string        `yaml:"host_port"`
	URL            string        `yaml:"url"`
	ConnectTimeout *TimeInterval `yaml:"connect_timeout"`
	Tag            *Tag          `yaml:"tag"`
	GRPCCompressor string        `yaml:"grpc_compressor"`
}
