package client

import (
	"encoding/json"
	"fmt"
	"time"
)

type ConfigParser interface {
	Decode(data []byte, config EtcdConfig) error
}

type defaultParser struct {
}

func (p *defaultParser) Decode(data []byte, config EtcdConfig) error {
	return json.Unmarshal(data, &config)
}

type Config interface {
	String() string
}

type EtcdConfig struct {
	ClientBasicInfo *EndpointBasicInfo `mapstructure:"ClientBasicInfo"`
	HostPorts       *string            `mapstructure:"HostPorts"`
	DestService     *string            `mapstructure:"DestService"`
	Protocol        *string            `mapstructure:"Protocol"`
	Connection      *Connection        `mapstructure:"Connection"`
	MyConfig        Config             `mapstructure:"MyConfig"`
}

func (c *EtcdConfig) String() string {
	baseInfo := "nil"
	if c.MyConfig != nil {
		baseInfo = c.MyConfig.String()
	}
	return fmt.Sprintf("ClientBasicInfo: %v\n"+
		" HostPorts: %s\n"+
		" DestService: %s\n"+
		" Protocol: %s\n"+
		" Connection: %v\n"+
		" MyConfig: %s\n",
		*c.ClientBasicInfo, *c.HostPorts, *c.DestService, *c.Protocol, *c.Connection, baseInfo)
}

type EndpointBasicInfo struct {
	ServiceName string            `mapstructure:"ServiceName"`
	Method      string            `mapstructure:"Method"`
	Tags        map[string]string `mapstructure:"Tags"`
}

type IdleConfig struct {
	MinIdlePerAddress int           `mapstructure:"MinIdlePerAddress"`
	MaxIdlePerAddress int           `mapstructure:"MaxIdlePerAddress"`
	MaxIdleGlobal     int           `mapstructure:"MaxIdleGlobal"`
	MaxIdleTimeout    time.Duration `mapstructure:"MaxIdleTimeout"`
}
type MuxConnection struct {
	ConnNum int `json:"connNum"`
}

type Connection struct {
	Method         string        `mapstructure:"Method"`
	LongConnection IdleConfig    `mapstructure:"LongConnection"`
	MuxConnection  MuxConnection `mapstructure:"MuxConnection"`
}
