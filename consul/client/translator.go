package client

import (
	"fmt"
	"github.com/Printemps417/optionloader/utils"
	kitexclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
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

func basicInfoTranslator(config *EtcdConfig) ([]kitexclient.Option, error) {
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
func protocolTranslator(config *EtcdConfig) ([]kitexclient.Option, error) {
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
func destServiceTranslator(config *EtcdConfig) ([]kitexclient.Option, error) {
	c := config.DestService
	if c == nil {
		return nil, nil
	}
	var res []kitexclient.Option
	res = append(res, kitexclient.WithDestService(*c))
	return res, nil
}

func hostPortsTranslator(config *EtcdConfig) ([]kitexclient.Option, error) {
	c := config.HostPorts
	if c == nil {
		return nil, nil
	}
	var res []kitexclient.Option
	res = append(res, kitexclient.WithHostPorts(c...))
	return res, nil
}
func connectionTranslator(config *EtcdConfig) ([]kitexclient.Option, error) {
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
