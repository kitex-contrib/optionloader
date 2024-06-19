package server

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"strings"
)

type ConfigParser interface {
	Decode(configType ConfigType, data []byte, config *ConsulConfig) error
}

type defaultParser struct {
}

func (p *defaultParser) Decode(configType ConfigType, data []byte, config *ConsulConfig) error {
	switch configType {
	case JSON:
		return json.Unmarshal(data, config)
	case YAML:
		return yaml.Unmarshal(data, config)
	default:
		return fmt.Errorf("unsupported config data type %s", configType)
	}
}

type Config interface {
	String() string
}

type ConsulConfig struct {
	ServerBasicInfo *EndpointBasicInfo `mapstructure:"ServerBasicInfo"`
	ServiceAddr     []Addr             `mapstructure:"ServiceAddr"`
	MuxTransport    *bool              `mapstructure:"MuxTransport"`
	MyConfig        Config             `mapstructure:"MyConfig"`
}

func (c *ConsulConfig) String() string {
	var builder strings.Builder
	if c.ServerBasicInfo != nil {
		builder.WriteString(fmt.Sprintf("ServerBasicInfo: %v\n", *c.ServerBasicInfo))
	}
	if c.ServiceAddr != nil {
		builder.WriteString(fmt.Sprintf("ServiceAddr: %v\n", c.ServiceAddr))
	}
	if c.MuxTransport != nil {
		builder.WriteString(fmt.Sprintf("MuxTransport: %v\n", *c.MuxTransport))
	}
	if c.MyConfig != nil {
		builder.WriteString(c.MyConfig.String())
	}
	return builder.String()
}

type EndpointBasicInfo struct {
	ServiceName string            `mapstructure:"ServiceName"`
	Method      string            `mapstructure:"Method"`
	Tags        map[string]string `mapstructure:"Tags"`
}

type Addr struct {
	Network string `mapstructure:"network"`
	Address string `mapstructure:"address"`
}
