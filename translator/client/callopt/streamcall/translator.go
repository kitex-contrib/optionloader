package streamcall

import (
	"github.com/cloudwego/kitex/client/callopt/streamcall"
	"github.com/kitex-contrib/optionloader/config"
)

func HostPortTranslator(config *config.StreamCallConfig) streamcall.Option {
	return streamcall.WithHostPort(config.HostPort)
}

func URLTranslator(config *config.StreamCallConfig) streamcall.Option {
	return streamcall.WithURL(config.URL)
}

func ConnectTimeoutTranslator(config *config.StreamCallConfig) streamcall.Option {
	return streamcall.WithConnectTimeout(config.ConnectTimeout.Transform())
}

func TagTranslator(config *config.StreamCallConfig) streamcall.Option {
	return streamcall.WithTag(config.Tag.Key, config.Tag.Value)
}

func GRPCCompressorTranslator(config *config.StreamCallConfig) streamcall.Option {
	return streamcall.WithGRPCCompressor(config.GRPCCompressor)
}
