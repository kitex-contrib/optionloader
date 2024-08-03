package callopt

import (
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/kitex-contrib/optionloader/config"
)

func HostPortTranslator(config *config.CalloptConfig) callopt.Option {
	return callopt.WithHostPort(config.HostPort)
}

func URLTranslator(config *config.CalloptConfig) callopt.Option {
	return callopt.WithURL(config.URL)
}

func HTTPHostTranslator(config *config.CalloptConfig) callopt.Option {
	return callopt.WithHTTPHost(config.HTTPHost)
}

func RPCTimeoutTranslator(config *config.CalloptConfig) callopt.Option {
	return callopt.WithRPCTimeout(config.RPCTimeout.Transform())
}

func ConnectTimeoutTranslator(config *config.CalloptConfig) callopt.Option {
	return callopt.WithConnectTimeout(config.ConnectTimeout.Transform())
}

func TagTranslator(config *config.CalloptConfig) callopt.Option {
	return callopt.WithTag(config.Tag.Key, config.Tag.Value)
}

func GRPCCompressorTranslator(config *config.CalloptConfig) callopt.Option {
	return callopt.WithGRPCCompressor(config.GRPCCompressor)
}
