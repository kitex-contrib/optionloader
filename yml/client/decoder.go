package client

import (
	"fmt"
	"github.com/cloudwego/kitex/pkg/circuitbreak"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpctimeout"
	"time"
)

// EndpointBasicInfo should be immutable after created.
type EndpointBasicInfo struct {
	ServiceName string            `yaml:"ServiceName"`
	Method      string            `yaml:"Method"`
	Tags        map[string]string `yaml:"Tags"`
}

type port string

type protocol string
type DestService string

// IdleConfig contains idle configuration for long-connection pool.
type IdleConfig struct {
	MinIdlePerAddress int           `yaml:"MinIdlePerAddress"`
	MaxIdlePerAddress int           `yaml:"MaxIdlePerAddress"`
	MaxIdleGlobal     int           `yaml:"MaxIdleGlobal"`
	MaxIdleTimeout    time.Duration `yaml:"MaxIdleTimeout"`
}
type MuxConnection struct {
	ConnNum int `yaml:"connNum"`
}

type Connection struct {
	Method         string        `yaml:"method"`
	LongConnection IdleConfig    `yaml:"LongConnection"`
	MuxConnection  MuxConnection `yaml:"MuxConnection"`
}

// RPCTimeout
type Timeout rpctimeout.RPCTimeout

// RetryPolicy
type RetryPolicy retry.Policy

// CircuitBreaker
type Circuitbreaker circuitbreak.CBConfig
type YMLConfig struct {
	ClientBasicInfo EndpointBasicInfo `yaml:"ClientBasicInfo"`
	HostPorts       int               `yaml:"HostPorts"`
	DestService     string            `yaml:"DestService"`
	Protocol        string            `yaml:"Protocol"`
	Connection      Connection        `yaml:"Connection"`
}

func (c YMLConfig) String() string {
	return fmt.Sprintf("ClientBasicInfo: %v\n"+
		" HostPorts: %d\n"+
		" DestService: %s\n"+
		" Protocol: %s\n"+
		" Connection: %v\n",
		c.ClientBasicInfo, c.HostPorts, c.DestService, c.Protocol, c.Connection)
}
