// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"encoding/json"
	"fmt"
	"github.com/cloudwego/kitex/pkg/retry"
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
	ClientBasicInfo   *EndpointBasicInfo       `mapstructure:"ClientBasicInfo"`
	HostPorts         []string                 `mapstructure:"HostPorts"`
	DestService       *string                  `mapstructure:"DestService"`
	Protocol          *string                  `mapstructure:"Protocol"`
	MyConfig          Config                   `mapstructure:"MyConfig"`
	Connection        *Connection              `mapstructure:"Connection"`
	FailureRetry      *retry.FailurePolicy     `mapstructure:"FailureRetry"`
	ShouldResultRetry *retry.ShouldResultRetry `mapstructure:"-"`
}

func (c *ConsulConfig) String() string {
	var builder strings.Builder

	if c.ClientBasicInfo != nil {
		builder.WriteString(fmt.Sprintf("ClientBasicInfo: %v\n", *c.ClientBasicInfo))
	}

	if c.HostPorts != nil {
		builder.WriteString(fmt.Sprintf("HostPorts: %v\n", c.HostPorts))
	}

	if c.DestService != nil {
		builder.WriteString(fmt.Sprintf("DestService: %v\n", *c.DestService))
	}

	if c.Protocol != nil {
		builder.WriteString(fmt.Sprintf("Protocol: %v\n", *c.Protocol))
	}

	if c.Connection != nil {
		builder.WriteString(fmt.Sprintf("Connection: %v\n", *c.Connection))
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

type IdleConfig struct {
	MinIdlePerAddress int    `mapstructure:"MinIdlePerAddress"`
	MaxIdlePerAddress int    `mapstructure:"MaxIdlePerAddress"`
	MaxIdleGlobal     int    `mapstructure:"MaxIdleGlobal"`
	MaxIdleTimeout    string `mapstructure:"MaxIdleTimeout"`
}
type MuxConnection struct {
	ConnNum int `mapstructure:"ConnNum"`
}

type Connection struct {
	Method         string        `mapstructure:"Method"`
	LongConnection IdleConfig    `mapstructure:"LongConnection"`
	MuxConnection  MuxConnection `mapstructure:"MuxConnection"`
}
