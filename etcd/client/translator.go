package client

import (
	"fmt"
	kitexclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/transport"
	"strconv"
	"strings"
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

func basicinfoTranslator(config *EtcdConfig) ([]kitexclient.Option, error) {
	c := config.ClientBasicInfo
	res := []kitexclient.Option{}
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
	res := []kitexclient.Option{}
	protocol, ok := protocolMap[*c]
	if !ok {
		return nil, fmt.Errorf("unknown protocol: %s", *c)
	}
	res = append(res, kitexclient.WithTransportProtocol(transport.Protocol(protocol)))

	return res, nil
}
func DestServiceTranslator(config *EtcdConfig) ([]kitexclient.Option, error) {
	c := config.DestService
	res := []kitexclient.Option{}
	res = append(res, kitexclient.WithDestService(*c))
	return res, nil
}

func HostPortsTranslator(config *EtcdConfig) ([]kitexclient.Option, error) {
	c := config.HostPorts
	res := []kitexclient.Option{}

	ports := strings.Split(*c, ",")
	for _, port := range ports {
		hostPort := strings.Split(port, ":")
		portNum, err := strconv.Atoi(hostPort[len(hostPort)-1])
		if err != nil || portNum < 1 || portNum > 65535 {
			return nil, fmt.Errorf("invalid port number: %s", port)
		}
	}

	res = append(res, kitexclient.WithHostPorts(*c))
	return res, nil
}
func connectionTranslator(config *EtcdConfig) ([]kitexclient.Option, error) {
	c := config.Connection
	res := []kitexclient.Option{}

	switch c.Method {
	case "ShortConnection":
		res = append(res, kitexclient.WithShortConnection())
	case "LongConnection":
		idleConfig := connpool.IdleConfig{
			MinIdlePerAddress: c.LongConnection.MinIdlePerAddress,
			MaxIdlePerAddress: c.LongConnection.MaxIdlePerAddress,
			MaxIdleGlobal:     c.LongConnection.MaxIdleGlobal,
			MaxIdleTimeout:    c.LongConnection.MaxIdleTimeout,
		}
		res = append(res, kitexclient.WithLongConnection(idleConfig))
	case "MuxConnection":
		res = append(res, kitexclient.WithMuxConnection(c.MuxConnection.ConnNum))
	default:
		return nil, fmt.Errorf("unsupported connection method: %s", c.Method)
	}

	return res, nil
}
