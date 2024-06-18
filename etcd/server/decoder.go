package etcdserver

import (
	"encoding/json"
	"fmt"
)

type ConfigParser interface {
	Decode(data []byte, config *EtcdConfig) error
}

type defaultParser struct {
}

func (p *defaultParser) Decode(data []byte, config *EtcdConfig) error {
	return json.Unmarshal(data, config)
}

type Config interface {
	String() string
}

type EtcdConfig struct {
	ServerBasicInfo *EndpointBasicInfo `mapstructure:"ServerBasicInfo"`
	ServiceAddr     []Addr             `mapstructure:"ServiceAddr"`
	MuxTransport    *bool              `mapstructure:"MuxTransport"`
	MyConfig        Config             `mapstructure:"MyConfig"`
}

func (c *EtcdConfig) String() string {
	baseInfo := "nil"
	if c.MyConfig != nil {
		baseInfo = c.MyConfig.String()
	}
	return fmt.Sprintf(
		" ServerBasicInfo: %s\n"+
			" ServiceAddr: %s\n"+
			" MuxTransport: %v\n"+
			" MyConfig: %s\n",
		*c.ServerBasicInfo, c.ServiceAddr, *c.MuxTransport, baseInfo)
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
