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

func basicinfoTranslator(p interface{}) ([]kitexclient.Option, error) {
	res := []kitexclient.Option{}
	switch v := p.(type) {
	case rpcinfo.EndpointBasicInfo:
		res = append(res, kitexclient.WithClientBasicInfo(&v))
	case EndpointBasicInfo: // Assuming EndpointBasicInfo is your custom type
		rpcInfo := rpcinfo.EndpointBasicInfo{
			ServiceName: v.ServiceName,
			Method:      v.Method,
			Tags:        v.Tags,
		}
		res = append(res, kitexclient.WithClientBasicInfo(&rpcInfo))
	default:
		return nil, fmt.Errorf("unsupported type: %T", p)
	}
	return res, nil
}
func protocolTranslator(p interface{}) ([]kitexclient.Option, error) {
	res := []kitexclient.Option{}
	switch v := p.(type) {
	case string:
		protocol, ok := protocolMap[v]
		if !ok {
			return nil, fmt.Errorf("unknown protocol: %s", v)
		}
		res = append(res, kitexclient.WithTransportProtocol(transport.Protocol(protocol)))
	default:
		return nil, fmt.Errorf("unsupported type: %T", p)
	}
	return res, nil
}
func DestServiceTranslator(p interface{}) ([]kitexclient.Option, error) {
	res := []kitexclient.Option{}
	str, ok := p.(string)
	if !ok {
		return nil, fmt.Errorf("DestService should be a string")
	}
	res = append(res, kitexclient.WithDestService(str))
	return res, nil
}

func HostPortsTranslator(p interface{}) ([]kitexclient.Option, error) {
	res := []kitexclient.Option{}
	str, _ := p.(string)

	// Split the string by comma for multiple ports
	ports := strings.Split(str, ",")
	for _, port := range ports {
		// Split the string by colon for host:port format
		hostPort := strings.Split(port, ":")
		portNum, err := strconv.Atoi(hostPort[len(hostPort)-1])
		if err != nil || portNum < 1 || portNum > 65535 {
			return nil, fmt.Errorf("invalid port number: %s", port)
		}
	}

	res = append(res, kitexclient.WithHostPorts(str))
	return res, nil
}
func connectionTranslator(p interface{}) ([]kitexclient.Option, error) {
	res := []kitexclient.Option{}
	conn, ok := p.(Connection)
	if !ok {
		return nil, fmt.Errorf("Connection should be a Connection type")
	}

	switch conn.Method {
	case "ShortConnection":
		res = append(res, kitexclient.WithShortConnection())
	case "LongConnection":
		idleConfig := connpool.IdleConfig{
			MinIdlePerAddress: conn.LongConnection.MinIdlePerAddress,
			MaxIdlePerAddress: conn.LongConnection.MaxIdlePerAddress,
			MaxIdleGlobal:     conn.LongConnection.MaxIdleGlobal,
			MaxIdleTimeout:    conn.LongConnection.MaxIdleTimeout,
		}
		res = append(res, kitexclient.WithLongConnection(idleConfig))
	case "MuxConnection":
		res = append(res, kitexclient.WithMuxConnection(conn.MuxConnection.ConnNum))
	default:
		return nil, fmt.Errorf("unsupported connection method: %s", conn.Method)
	}

	return res, nil
}
