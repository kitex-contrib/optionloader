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
	"errors"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/stats"
	kitexserver "github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/optionloader/utils"
	"net"
)

type Translator func(config ConsulConfig) ([]kitexserver.Option, error)

func basicInfoTranslator(config ConsulConfig) ([]kitexserver.Option, error) {
	c := config.ServerBasicInfo
	if c == nil {
		return nil, nil
	}
	var res []kitexserver.Option
	rpcInfo := rpcinfo.EndpointBasicInfo{
		ServiceName: c.ServiceName,
		Method:      c.Method,
		Tags:        c.Tags,
	}
	res = append(res, kitexserver.WithServerBasicInfo(&rpcInfo))
	return res, nil
}
func serviceAddrTranslator(config ConsulConfig) ([]kitexserver.Option, error) {
	c := config.ServiceAddr
	if c == nil {
		return nil, nil
	}
	var res []kitexserver.Option
	for _, addr := range c {
		network := addr.Network
		address := addr.Address
		var addr net.Addr
		var err error
		switch network {
		case "tcp", "tcp4", "tcp6":
			addr, err = net.ResolveTCPAddr(network, address)
		case "ip", "ip4", "ip6":
			addr, err = net.ResolveIPAddr(network, address)
		case "udp", "udp4", "udp6":
			addr, err = net.ResolveUDPAddr(network, address)
		case "unix", "unixgram", "unixpacket":
			addr, err = net.ResolveUnixAddr(network, address)
		default:
			err = errors.New("unknown network: " + network)
		}
		if err != nil {
			return nil, err
		}
		res = append(res, kitexserver.WithServiceAddr(addr))
	}
	return res, nil
}
func muxTransportTranslator(config ConsulConfig) ([]kitexserver.Option, error) {
	c := config.MuxTransport
	if c == nil {
		return nil, nil
	}
	var res []kitexserver.Option
	if *c {
		res = append(res, kitexserver.WithMuxTransport())
	}
	return res, nil
}
func readWriteTimeoutTranslator(config ConsulConfig) ([]kitexserver.Option, error) {
	c := config.ReadWriteTimeout
	if c == nil {
		return nil, nil
	}
	var res []kitexserver.Option
	timeout, err := utils.ParseDuration(*c)
	if err != nil {
		return nil, err
	}
	res = append(res, kitexserver.WithReadWriteTimeout(timeout))
	return res, nil
}
func exitWaitTimeTranslator(config ConsulConfig) ([]kitexserver.Option, error) {
	c := config.ExitWaitTime
	if c == nil {
		return nil, nil
	}
	var res []kitexserver.Option
	timeout, err := utils.ParseDuration(*c)
	if err != nil {
		return nil, err
	}
	res = append(res, kitexserver.WithExitWaitTime(timeout))
	return res, nil
}
func maxConnIdleTimeTranslator(config ConsulConfig) ([]kitexserver.Option, error) {
	c := config.MaxConnIdleTime
	if c == nil {
		return nil, nil
	}
	var res []kitexserver.Option
	timeout, err := utils.ParseDuration(*c)
	if err != nil {
		return nil, err
	}
	res = append(res, kitexserver.WithMaxConnIdleTime(timeout))
	return res, nil
}
func statsLevelTranslator(config ConsulConfig) ([]kitexserver.Option, error) {
	c := config.StatsLevel
	if c == nil {
		return nil, nil
	}
	var res []kitexserver.Option
	res = append(res, kitexserver.WithStatsLevel(stats.Level(*c)))
	return res, nil
}
func grpcTranslator(config ConsulConfig) ([]kitexserver.Option, error) {
	c := config.GRPC
	if c == nil {
		return nil, nil
	}
	var res []kitexserver.Option
	if c.WriteBufferSize != nil {
		res = append(res, kitexserver.WithGRPCWriteBufferSize(*c.WriteBufferSize))
	}
	if c.ReadBufferSize != nil {
		res = append(res, kitexserver.WithGRPCReadBufferSize(*c.ReadBufferSize))
	}
	if c.InitialWindowSize != nil {
		res = append(res, kitexserver.WithGRPCInitialWindowSize(*c.InitialWindowSize))
	}
	if c.InitialConnWindowSize != nil {
		res = append(res, kitexserver.WithGRPCInitialConnWindowSize(*c.InitialConnWindowSize))
	}
	if c.KeepaliveParams != nil {
		k := c.KeepaliveParams
		maxConnectionIdle, err := utils.ParseDuration(k.MaxConnectionIdle)
		if err != nil {
			return nil, err
		}
		maxConnectionAge, err := utils.ParseDuration(k.MaxConnectionAge)
		if err != nil {
			return nil, err
		}
		maxConnectionAgeGrace, err := utils.ParseDuration(k.MaxConnectionAgeGrace)
		if err != nil {
			return nil, err
		}
		time, err := utils.ParseDuration(k.Time)
		if err != nil {
			return nil, err
		}
		timeout, err := utils.ParseDuration(k.Timeout)
		if err != nil {
			return nil, err
		}
		serverKeepalive := grpc.ServerKeepalive{
			MaxConnectionIdle:     maxConnectionIdle,
			MaxConnectionAge:      maxConnectionAge,
			MaxConnectionAgeGrace: maxConnectionAgeGrace,
			Time:                  time,
			Timeout:               timeout,
		}
		res = append(res, kitexserver.WithGRPCKeepaliveParams(serverKeepalive))
	}
	if c.KeepaliveEnforcementPolicy != nil {
		k := c.KeepaliveEnforcementPolicy
		minTimem, err := utils.ParseDuration(k.MinTime)
		if err != nil {
			return nil, err
		}
		enforcementPolicy := grpc.EnforcementPolicy{
			MinTime:             minTimem,
			PermitWithoutStream: k.PermitWithoutStream,
		}
		res = append(res, kitexserver.WithGRPCKeepaliveEnforcementPolicy(enforcementPolicy))
	}
	if c.MaxConcurrentStreams != nil {
		res = append(res, kitexserver.WithGRPCMaxConcurrentStreams(*c.MaxConcurrentStreams))
	}
	if c.MaxHeaderListSize != nil {
		res = append(res, kitexserver.WithGRPCMaxHeaderListSize(*c.MaxHeaderListSize))
	}
	return res, nil
}
