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
)

type Config interface {
}

type ConsulConfig struct {
	ServerBasicInfo  *EndpointBasicInfo `mapstructure:"ServerBasicInfo"`
	ServiceAddr      []Addr             `mapstructure:"ServiceAddr"`
	MuxTransport     *bool              `mapstructure:"MuxTransport"`
	ReadWriteTimeout *string            `mapstructure:"ReadWriteTimeout"`
	ExitWaitTime     *string            `mapstructure:"ExitWaitTime"`
	MaxConnIdleTime  *string            `mapstructure:"MaxConnIdleTime"`
	StatsLevel       *int               `mapstructure:"StatsLevel"`
	GRPC             *Grpc              `mapstructure:"GRPC"`
	MyConfig         Config             `mapstructure:"MyConfig"`
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

type Grpc struct {
	WriteBufferSize            *uint32            `mapstructure:"WriteBufferSize"`
	ReadBufferSize             *uint32            `mapstructure:"ReadBufferSize"`
	InitialWindowSize          *uint32            `mapstructure:"InitialWindowSize"`
	InitialConnWindowSize      *uint32            `mapstructure:"InitialConnWindowSize"`
	KeepaliveParams            *ServerKeepalive   `mapstructure:"KeepaliveParams"`
	KeepaliveEnforcementPolicy *EnforcementPolicy `mapstructure:"KeepaliveEnforcementPolicy"`
	MaxConcurrentStreams       *uint32            `mapstructure:"MaxConcurrentStreams"`
	MaxHeaderListSize          *uint32            `mapstructure:"MaxHeaderListSize"`
}

type ServerKeepalive struct {
	MaxConnectionIdle     string `mapstructure:"MaxConnectionIdle"`
	MaxConnectionAge      string `mapstructure:"MaxConnectionAge"`
	MaxConnectionAgeGrace string `mapstructure:"MaxConnectionAgeGrace"`
	Time                  string `mapstructure:"Time"`
	Timeout               string `mapstructure:"Timeout"`
}

type EnforcementPolicy struct {
	MinTime             string `mapstructure:"MinTime"`
	PermitWithoutStream bool   `mapstructure:"PermitWithoutStream"`
}
