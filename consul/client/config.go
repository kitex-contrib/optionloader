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
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/stats"
)

type Config interface {
}

type ConsulConfig struct {
	ClientBasicInfo   *EndpointBasicInfo       `mapstructure:"ClientBasicInfo"`
	HostPorts         []string                 `mapstructure:"HostPorts"`
	DestService       *string                  `mapstructure:"DestService"`
	Protocol          *string                  `mapstructure:"Protocol"`
	Connection        *Connection              `mapstructure:"Connection"`
	FailureRetry      *FailurePolicy           `mapstructure:"FailureRetry"`
	ShouldResultRetry *retry.ShouldResultRetry `mapstructure:"-"`
	BackupRequest     *BackupPolicy            `mapstructure:"BackupRequest"`
	RPCTimeout        *string                  `mapstructure:"RPCTimeout"`
	ConnectionTimeout *string                  `mapstructure:"ConnectionTimeout"`
	Tags              []Tag                    `mapstructure:"Tags"`
	StatsLevel        *stats.Level             `mapstructure:"StatsLevel"`
	GRPC              *Grpc                    `mapstructure:"GRPC"`
	CallOpt           *CallOpt                 `mapstructure:"CallOpt"`
	Stream            *StreamConfig            `mapstructure:"Stream"`
	StreamCallOpt     *StreamCallOpt           `mapstructure:"StreamCallOpt"`
	MyConfig          Config                   `mapstructure:"MyConfig"`
}

func (c *ConsulConfig) String() string {
	marshal, err := json.Marshal(c)
	if err != nil {
		return ""
	}
	return string(marshal)
}

type BackOffType string
type BackOffCfgKey string
type Type int

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

type FailurePolicy struct {
	StopPolicy        StopPolicy               `mapstructure:"StopPolicy"`
	BackOffPolicy     *BackOffPolicy           `mapstructure:"BackOffPolicy"`
	RetrySameNode     bool                     `mapstructure:"RetrySameNode"`
	ShouldResultRetry *retry.ShouldResultRetry `mapstructure:"-"`
	Extra             string                   `mapstructure:"Extra"`
}

type StopPolicy struct {
	MaxRetryTimes    int      `mapstructure:"MaxRetryTimes"`
	MaxDurationMS    uint32   `mapstructure:"MaxDurationMS"`
	DisableChainStop bool     `mapstructure:"DisableChainStop"`
	DDLStop          bool     `mapstructure:"DDLStop"`
	CBPolicy         CBPolicy `mapstructure:"CBPolicy"`
}

type CBPolicy struct {
	ErrorRate float64 `mapstructure:"ErrorRate"`
}

type BackOffPolicy struct {
	BackOffType BackOffType               `mapstructure:"BackOffType"`
	CfgItems    map[BackOffCfgKey]float64 `mapstructure:"CfgItems"`
}

type BackupPolicy struct {
	RetryDelayMS  uint32     `mapstructure:"RetryDelayMS"`
	StopPolicy    StopPolicy `mapstructure:"StopPolicy"`
	RetrySameNode bool       `mapstructure:"RetrySameNode"`
}

type Tag struct {
	Key   string `mapstructure:"Key"`
	Value string `mapstructure:"Value"`
}

type Grpc struct {
	GRPCConnPoolSize          *uint32              `mapstructure:"GRPCConnPoolSize"`
	GRPCWriteBufferSize       *uint32              `mapstructure:"GRPCWriteBufferSize"`
	GRPCReadBufferSize        *uint32              `mapstructure:"GRPCReadBufferSize"`
	GRPCInitialWindowSize     *uint32              `mapstructure:"GRPCInitialWindowSize"`
	GRPCInitialConnWindowSize *uint32              `mapstructure:"GRPCInitialConnWindowSize"`
	GRPCMaxHeaderListSize     *uint32              `mapstructure:"GRPCMaxHeaderListSize"`
	GRPCKeepaliveParams       *GRPCClientKeepalive `mapstructure:"GRPCKeepaliveParams"`
}

type GRPCClientKeepalive struct {
	Time                string `mapstructure:"Time"`
	Timeout             string `mapstructure:"Timeout"`
	PermitWithoutStream bool   `mapstructure:"PermitWithoutStream"`
}

type CallOpt struct {
	HostPorts         *map[string]string `mapstructure:"HostPorts"`
	Urls              *map[string]string `mapstructure:"Urls"`
	Tags              *map[string]Tag    `mapstructure:"Tags"`
	RPCTimeout        *string            `mapstructure:"RPCTimeout"`
	ConnectionTimeout *string            `mapstructure:"ConnectionTimeout"`
	HTTPHost          *string            `mapstructure:"HTTPHost"`
	RetryPolicy       *Policy            `mapstructure:"RetryPolicy"`
	CompressorName    *string            `mapstructure:"CompressorName"`
}

type StreamConfig struct {
	ClientBasicInfo   *EndpointBasicInfo `mapstructure:"ClientBasicInfo"`
	HostPorts         []string           `mapstructure:"HostPorts"`
	DestService       *string            `mapstructure:"DestService"`
	ConnectionTimeout *string            `mapstructure:"ConnectionTimeout"`
	Tags              []Tag              `mapstructure:"Tags"`
	StatsLevel        *stats.Level       `mapstructure:"StatsLevel"`
	GRPC              *Grpc              `mapstructure:"GRPC"`
}

type StreamCallOpt struct {
	HostPorts         *map[string]string `mapstructure:"HostPorts"`
	Urls              *map[string]string `mapstructure:"Urls"`
	Tags              *map[string]Tag    `mapstructure:"Tags"`
	ConnectionTimeout *string            `mapstructure:"ConnectionTimeout"`
	CompressorName    *string            `mapstructure:"CompressorName"`
}

type Policy struct {
	Enable        bool           `mapstructure:"Enable"`
	Type          Type           `mapstructure:"Type"`
	FailurePolicy *FailurePolicy `mapstructure:"FailurePolicy"`
	BackupPolicy  *BackupPolicy  `mapstructure:"BackupPolicy"`
}
