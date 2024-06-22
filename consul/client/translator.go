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
	"fmt"
	kitexclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/optionloader/utils"
	"github.com/mitchellh/mapstructure"
)

// Protocol indicates the transport protocol.
type Protocol int

// Predefined transport protocols.
const (
	PurePayload Protocol = 0

	TTHeader Protocol = 1 << iota
	Framed
	HTTP
	GRPC
	HESSIAN2

	TTHeaderFramed = TTHeader | Framed
)

var protocolMap = map[string]Protocol{
	"PurePayload":    PurePayload,
	"TTHeader":       TTHeader,
	"Framed":         Framed,
	"HTTP":           HTTP,
	"GRPC":           GRPC,
	"HESSIAN2":       HESSIAN2,
	"TTHeaderFramed": TTHeaderFramed,
}

func basicInfoTranslator(config *ConsulConfig) ([]kitexclient.Option, error) {
	c := config.ClientBasicInfo
	if c == nil {
		return nil, nil
	}
	var res []kitexclient.Option
	rpcInfo := rpcinfo.EndpointBasicInfo{
		ServiceName: c.ServiceName,
		Method:      c.Method,
		Tags:        c.Tags,
	}
	res = append(res, kitexclient.WithClientBasicInfo(&rpcInfo))
	return res, nil
}
func protocolTranslator(config *ConsulConfig) ([]kitexclient.Option, error) {
	c := config.Protocol
	if c == nil {
		return nil, nil
	}
	var res []kitexclient.Option
	protocol, ok := protocolMap[*c]
	if !ok {
		return nil, fmt.Errorf("unknown protocol: %s", *c)
	}
	res = append(res, kitexclient.WithTransportProtocol(transport.Protocol(protocol)))

	return res, nil
}
func destServiceTranslator(config *ConsulConfig) ([]kitexclient.Option, error) {
	c := config.DestService
	if c == nil {
		return nil, nil
	}
	var res []kitexclient.Option
	res = append(res, kitexclient.WithDestService(*c))
	return res, nil
}

func hostPortsTranslator(config *ConsulConfig) ([]kitexclient.Option, error) {
	c := config.HostPorts
	if c == nil {
		return nil, nil
	}
	var res []kitexclient.Option
	res = append(res, kitexclient.WithHostPorts(c...))
	return res, nil
}
func connectionTranslator(config *ConsulConfig) ([]kitexclient.Option, error) {
	c := config.Connection
	if c == nil {
		return nil, nil
	}
	var res []kitexclient.Option

	switch c.Method {
	case "ShortConnection":
		res = append(res, kitexclient.WithShortConnection())
	case "LongConnection":
		MaxIdleTimeout, err := utils.ParseDuration(c.LongConnection.MaxIdleTimeout)
		if err != nil {
			return nil, err
		}
		idleConfig := connpool.IdleConfig{
			MinIdlePerAddress: c.LongConnection.MinIdlePerAddress,
			MaxIdlePerAddress: c.LongConnection.MaxIdlePerAddress,
			MaxIdleGlobal:     c.LongConnection.MaxIdleGlobal,
			MaxIdleTimeout:    MaxIdleTimeout,
		}
		res = append(res, kitexclient.WithLongConnection(idleConfig))
	case "MuxConnection":
		res = append(res, kitexclient.WithMuxConnection(c.MuxConnection.ConnNum))
	default:
		return nil, fmt.Errorf("unsupported connection method: %s", c.Method)
	}

	return res, nil
}
func failureRetryTranslator(config *ConsulConfig) ([]kitexclient.Option, error) {
	c := config.FailureRetry
	if c == nil {
		return nil, nil
	}
	var res []kitexclient.Option
	failurePolicy := &retry.FailurePolicy{}
	err := mapstructure.Decode(*c, failurePolicy)
	if err != nil {
		return nil, err
	}
	res = append(res, kitexclient.WithFailureRetry(failurePolicy))
	return res, nil
}
func specifiedResultRetryTranslator(config *ConsulConfig) ([]kitexclient.Option, error) {
	c := config.ShouldResultRetry
	if c == nil {
		return nil, nil
	}
	var res []kitexclient.Option
	res = append(res, kitexclient.WithSpecifiedResultRetry(c))
	return res, nil
}
func backupRequestTranslator(config *ConsulConfig) ([]kitexclient.Option, error) {
	c := config.BackupRequest
	if c == nil {
		return nil, nil
	}
	var res []kitexclient.Option
	backupPolicy := &retry.BackupPolicy{}
	err := mapstructure.Decode(*c, backupPolicy)
	if err != nil {
		return nil, err
	}
	res = append(res, kitexclient.WithBackupRequest(backupPolicy))
	return res, nil
}
func rpcTimeoutTranslator(config *ConsulConfig) ([]kitexclient.Option, error) {
	c := config.RPCTimeout
	if c == nil {
		return nil, nil
	}
	var res []kitexclient.Option
	rpcTimeout, err := utils.ParseDuration(*c)
	if err != nil {
		return nil, err
	}
	res = append(res, kitexclient.WithRPCTimeout(rpcTimeout))
	return res, nil
}
func connectionTimeoutTranslator(config *ConsulConfig) ([]kitexclient.Option, error) {
	c := config.ConnectionTimeout
	if c == nil {
		return nil, nil
	}
	var res []kitexclient.Option
	rpcTimeout, err := utils.ParseDuration(*c)
	if err != nil {
		return nil, err
	}
	res = append(res, kitexclient.WithConnectTimeout(rpcTimeout))
	return res, nil
}
func tagsTranslator(config *ConsulConfig) ([]kitexclient.Option, error) {
	c := config.Tags
	if c == nil {
		return nil, nil
	}
	var res []kitexclient.Option
	for _, tag := range c {
		res = append(res, kitexclient.WithTag(tag.Key, tag.Value))
	}
	return res, nil
}
func statsLevelTranslator(config *ConsulConfig) ([]kitexclient.Option, error) {
	c := config.StatsLevel
	if c == nil {
		return nil, nil
	}
	var res []kitexclient.Option
	res = append(res, kitexclient.WithStatsLevel(*c))
	return res, nil
}

func grpcTranslator(config *ConsulConfig) ([]kitexclient.Option, error) {
	c := config.GRPC
	if c == nil {
		return nil, nil
	}
	var res []kitexclient.Option
	if c.GRPCConnPoolSize != nil {
		res = append(res, kitexclient.WithGRPCConnPoolSize(*c.GRPCConnPoolSize))
	}
	if c.GRPCWriteBufferSize != nil {
		res = append(res, kitexclient.WithGRPCWriteBufferSize(*c.GRPCWriteBufferSize))
	}
	if c.GRPCReadBufferSize != nil {
		res = append(res, kitexclient.WithGRPCReadBufferSize(*c.GRPCReadBufferSize))
	}
	if c.GRPCInitialWindowSize != nil {
		res = append(res, kitexclient.WithGRPCInitialWindowSize(*c.GRPCInitialWindowSize))
	}
	if c.GRPCInitialConnWindowSize != nil {
		res = append(res, kitexclient.WithGRPCInitialConnWindowSize(*c.GRPCInitialConnWindowSize))
	}
	if c.GRPCMaxHeaderListSize != nil {
		res = append(res, kitexclient.WithGRPCMaxHeaderListSize(*c.GRPCMaxHeaderListSize))
	}
	if c.GRPCKeepaliveParams != nil {
		keepaliveTime, err := utils.ParseDuration(c.GRPCKeepaliveParams.Time)
		if err != nil {
			return nil, err
		}
		keepaliveTimeout, err := utils.ParseDuration(c.GRPCKeepaliveParams.Timeout)
		if err != nil {
			return nil, err
		}
		keepaliveParams := grpc.ClientKeepalive{
			Time:                keepaliveTime,
			Timeout:             keepaliveTimeout,
			PermitWithoutStream: c.GRPCKeepaliveParams.PermitWithoutStream,
		}
		res = append(res, kitexclient.WithGRPCKeepaliveParams(keepaliveParams))
	}
	return res, nil
}
