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
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	kitexserver "github.com/cloudwego/kitex/server"
	"net"
)

func basicInfoTranslator(config *EtcdConfig) ([]kitexserver.Option, error) {
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
func serviceAddrTranslator(config *EtcdConfig) ([]kitexserver.Option, error) {
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
func muxTransportTranslator(config *EtcdConfig) ([]kitexserver.Option, error) {
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
