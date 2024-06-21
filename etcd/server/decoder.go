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

package server

import (
	"encoding/json"
	"fmt"
	"strings"
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
