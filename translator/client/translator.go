package client

import (
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/connpool"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/stats"
	"github.com/cloudwego/kitex/transport"
	"github.com/kitex-contrib/optionloader/config"
)

// basic
func TransportProtocolTranslator(config *config.ClientConfig) client.Option {
	return client.WithTransportProtocol(transport.Protocol(config.TransportProtocol))
}

func DestServiceTranslator(config *config.ClientConfig) client.Option {
	return client.WithDestService(config.DestService)
}

func HostPortsTranslator(config *config.ClientConfig) client.Option {
	return client.WithHostPorts(config.HostPorts...)
}

func LongConnectionTranslator(config *config.ClientConfig) client.Option {
	cfg := connpool.IdleConfig{
		MinIdlePerAddress: config.LongConnection.MinIdlePerAddress,
		MaxIdlePerAddress: config.LongConnection.MaxIdlePerAddress,
		MaxIdleGlobal:     config.LongConnection.MaxIdleGlobal,
		MaxIdleTimeout:    config.LongConnection.MaxIdleTimeout.Transform(),
	}
	return client.WithLongConnection(cfg)
}

func MuxConnectionTranslator(config *config.ClientConfig) client.Option {
	return client.WithMuxConnection(config.MuxConnection)
}

func RPCTimeoutTranslator(config *config.ClientConfig) client.Option {
	return client.WithRPCTimeout(config.RPCTimeout.Transform())
}

func ConnectTimeoutTranslator(config *config.ClientConfig) client.Option {
	return client.WithConnectTimeout(config.ConnectTimeout.Transform())
}

func TagTranslator(config *config.ClientConfig) client.Option {
	return client.WithTag(config.Tag.Key, config.Tag.Value)
}

func StatsLevelTranslator(config *config.ClientConfig) client.Option {
	return client.WithStatsLevel(stats.Level(config.StatsLevel))
}

func ConnReporterEnabledTranslator(config *config.ClientConfig) client.Option {
	return client.WithConnReporterEnabled()
}

func GRPCConnPoolSizeTranslator(config *config.ClientConfig) client.Option {
	return client.WithGRPCConnPoolSize(config.GRPCConnPoolSize)
}

func GRPCWriteBufferSizeTranslator(config *config.ClientConfig) client.Option {
	return client.WithGRPCWriteBufferSize(config.GRPCWriteBufferSize)
}

func GRPCReadBufferSizeTranslator(config *config.ClientConfig) client.Option {
	return client.WithGRPCReadBufferSize(config.GRPCReadBufferSize)
}

func GRPCInitialWindowSizeTranslator(config *config.ClientConfig) client.Option {
	return client.WithGRPCInitialWindowSize(config.GRPCInitialWindowSize)
}

func GRPCInitialConnWindowSizeTranslator(config *config.ClientConfig) client.Option {
	return client.WithGRPCInitialConnWindowSize(config.GRPCInitialConnWindowSize)
}

func GRPCMaxHeaderListSizeTranslator(config *config.ClientConfig) client.Option {
	return client.WithGRPCMaxHeaderListSize(config.GRPCMaxHeaderListSize)
}

func GRPCKeepaliveParamsTranslator(config *config.ClientConfig) client.Option {
	kp := grpc.ClientKeepalive{
		Time:                config.GRPCKeepaliveParams.Time.Transform(),
		Timeout:             config.GRPCKeepaliveParams.Timeout.Transform(),
		PermitWithoutStream: config.GRPCKeepaliveParams.PermitWithoutStream,
	}
	return client.WithGRPCKeepaliveParams(kp)
}

// advanced
func HTTPConnectionTranslator(config *config.ClientConfig) client.Option {
	return client.WithHTTPConnection()
}

func ClientBasicInfoTranslator(config *config.ClientConfig) client.Option {
	ebi := &rpcinfo.EndpointBasicInfo{
		ServiceName: config.ClientBasicInfo.ServiceName,
		Method:      config.ClientBasicInfo.Method,
		Tags:        config.ClientBasicInfo.Tags,
	}
	return client.WithClientBasicInfo(ebi)
}
