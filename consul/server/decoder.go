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
	"gopkg.in/yaml.v3"
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
}

type ConsulConfig struct {
	ServerBasicInfo *EndpointBasicInfo `mapstructure:"ServerBasicInfo"`
	ServiceAddr     []Addr             `mapstructure:"ServiceAddr"`
	MuxTransport    *bool              `mapstructure:"MuxTransport"`
	MyConfig        Config             `mapstructure:"MyConfig"`
}

func (c *ConsulConfig) String() string {
	marshal, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(marshal)
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
