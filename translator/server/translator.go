package server

import (
	"github.com/cloudwego/kitex/pkg/remote/trans/nphttp2/grpc"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/pkg/stats"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/optionloader/config"
)

// basic
func MuxTransportTranslator(config *config.ServerConfig) server.Option {
	return server.WithMuxTransport()
}

func ReadWriteTimeoutTranslator(config *config.ServerConfig) server.Option {
	duration := config.ReadWriteTimeout.Transform()
	return server.WithReadWriteTimeout(duration)
}

func ExitWaitTimeTranslator(config *config.ServerConfig) server.Option {
	duration := config.ExitWaitTime.Transform()
	return server.WithExitWaitTime(duration)
}

func MaxConnIdleTimeTranslator(config *config.ServerConfig) server.Option {
	duration := config.MaxConnIdleTime.Transform()
	return server.WithMaxConnIdleTime(duration)
}

func StatsLevelTranslator(config *config.ServerConfig) server.Option {
	return server.WithStatsLevel(stats.Level(config.StatsLevel))
}

type netAddr struct {
	network string
	address string
}

func (s netAddr) Network() string { return s.network }

func (s netAddr) String() string { return s.address }

func ServiceAddrTranslator(config *config.ServerConfig) server.Option {
	return server.WithServiceAddr(netAddr{
		network: config.ServiceAddr.Network,
		address: config.ServiceAddr.Address,
	})
}

func GRPCWriteBufferSizeTranslator(config *config.ServerConfig) server.Option {
	return server.WithGRPCWriteBufferSize(config.GRPCWriteBufferSize)
}

func GRPCReadBufferSizeTranslator(config *config.ServerConfig) server.Option {
	return server.WithGRPCReadBufferSize(config.GRPCReadBufferSize)
}

func GRPCInitialWindowSizeTranslator(config *config.ServerConfig) server.Option {
	return server.WithGRPCInitialWindowSize(config.GRPCInitialWindowSize)
}

func GRPCInitialConnWindowSizeTranslator(config *config.ServerConfig) server.Option {
	return server.WithGRPCInitialConnWindowSize(config.GRPCInitialConnWindowSize)
}

func GRPCKeepaliveParamsTranslator(config *config.ServerConfig) server.Option {
	return server.WithGRPCKeepaliveParams(grpc.ServerKeepalive{
		MaxConnectionIdle:     config.GRPCKeepaliveParams.MaxConnectionIdle.Transform(),
		MaxConnectionAge:      config.GRPCKeepaliveParams.MaxConnectionAge.Transform(),
		MaxConnectionAgeGrace: config.GRPCKeepaliveParams.MaxConnectionAgeGrace.Transform(),
		Time:                  config.GRPCKeepaliveParams.Time.Transform(),
		Timeout:               config.GRPCKeepaliveParams.Timeout.Transform(),
	})
}

func GRPCKeepaliveEnforcementPolicyTranslator(config *config.ServerConfig) server.Option {
	return server.WithGRPCKeepaliveEnforcementPolicy(grpc.EnforcementPolicy{
		MinTime:             config.GRPCKeepaliveEnforcementPolicy.MinTime.Transform(),
		PermitWithoutStream: config.GRPCKeepaliveEnforcementPolicy.PermitWithoutStream,
	})
}

func GRPCMaxConcurrentStreamsTranslator(config *config.ServerConfig) server.Option {
	return server.WithGRPCMaxConcurrentStreams(config.GRPCMaxConcurrentStreams)
}

func GRPCMaxHeaderListSizeTranslator(config *config.ServerConfig) server.Option {
	return server.WithGRPCMaxHeaderListSize(config.GRPCMaxHeaderListSize)
}

func ContextBackupTranslator(config *config.ServerConfig) server.Option {
	return server.WithContextBackup(config.ContextBackup.Enable, config.ContextBackup.Async)
}

func RefuseTrafficWithoutServiceNameTranslator(config *config.ServerConfig) server.Option {
	return server.WithRefuseTrafficWithoutServiceName()
}

func EnableContextTimeoutTranslator(config *config.ServerConfig) server.Option {
	return server.WithEnableContextTimeout(config.EnableContextTimeout)
}

// advanced
func ServerBasicInfoTranslator(config *config.ServerConfig) server.Option {
	endpointBasicInfo := &rpcinfo.EndpointBasicInfo{
		ServiceName: config.ServerBasicInfo.ServiceName,
		Method:      config.ServerBasicInfo.Method,
		Tags:        config.ServerBasicInfo.Tags,
	}
	return server.WithServerBasicInfo(endpointBasicInfo)
}

func ReusePortTranslator(config *config.ServerConfig) server.Option {
	return server.WithReusePort(config.ReusePort)
}

// stream
func CompatibleMiddlewareForUnaryTranslator(config *config.ServerConfig) server.Option {
	return server.WithCompatibleMiddlewareForUnary()
}
