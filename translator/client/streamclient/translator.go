package streamclient

import (
	"github.com/cloudwego/kitex/client/streamclient"
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/stats"
	"github.com/kitex-contrib/optionloader/config"
)

// basic
func DestServiceTranslator(config *config.StreamClientConfig) streamclient.Option {
	return streamclient.WithDestService(config.DestService)
}

func HostPortsTranslator(config *config.StreamClientConfig) streamclient.Option {
	return streamclient.WithHostPorts(config.HostPorts...)
}

func ConnectTimeoutTranslator(config *config.StreamClientConfig) streamclient.Option {
	return streamclient.WithConnectTimeout(config.ConnectTimeout.Transform())
}

func TagTranslator(config *config.StreamClientConfig) streamclient.Option {
	return streamclient.WithTag(config.Tag.Key, config.Tag.Value)
}

func StatsLevelTranslator(config *config.StreamClientConfig) streamclient.Option {
	return streamclient.WithStatsLevel(stats.Level(config.StatsLevel))
}

func ConnReporterEnabledTranslator(config *config.StreamClientConfig) streamclient.Option {
	return streamclient.WithConnReporterEnabled()
}

func GRPCConnPoolSizeTranslator(config *config.StreamClientConfig) streamclient.Option {
	return streamclient.WithGRPCConnPoolSize(config.GRPCConnPoolSize)
}

func GRPCWriteBufferSizeTranslator(config *config.StreamClientConfig) streamclient.Option {
	return streamclient.WithGRPCWriteBufferSize(config.GRPCWriteBufferSize)
}

func GRPCReadBufferSizeTranslator(config *config.StreamClientConfig) streamclient.Option {
	return streamclient.WithGRPCReadBufferSize(config.GRPCReadBufferSize)
}

func GRPCInitialWindowSizeTranslator(config *config.StreamClientConfig) streamclient.Option {
	return streamclient.WithGRPCInitialWindowSize(config.GRPCInitialWindowSize)
}

func GRPCInitialConnWindowSizeTranslator(config *config.StreamClientConfig) streamclient.Option {
	return streamclient.WithGRPCInitialConnWindowSize(config.GRPCInitialConnWindowSize)
}

func GRPCMaxHeaderListSizeTranslator(config *config.StreamClientConfig) streamclient.Option {
	return streamclient.WithGRPCMaxHeaderListSize(config.GRPCMaxHeaderListSize)
}

func GRPCKeepaliveParamsTranslator(config *config.StreamClientConfig) streamclient.Option {
	kp := grpc.ClientKeepalive{
		Time:                config.GRPCKeepaliveParams.Time.Transform(),
		Timeout:             config.GRPCKeepaliveParams.Timeout.Transform(),
		PermitWithoutStream: config.GRPCKeepaliveParams.PermitWithoutStream,
	}
	return streamclient.WithGRPCKeepaliveParams(kp)
}

// advanced
func ClientBasicInfoTranslator(config *config.StreamClientConfig) streamclient.Option {
	ebi := &rpcinfo.EndpointBasicInfo{
		ServiceName: config.ClientBasicInfo.ServiceName,
		Method:      config.ClientBasicInfo.Method,
		Tags:        config.ClientBasicInfo.Tags,
	}
	return streamclient.WithClientBasicInfo(ebi)
}
