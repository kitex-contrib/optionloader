package client

import (
	"fmt"
	"github.com/cloudwego/kitex/client"
	"github.com/kitex-contrib/optionloader/config"
	translator "github.com/kitex-contrib/optionloader/translator/client"
)

type Translator func(config *config.ClientConfig) client.Option

type OptionLoader interface {
	// RegisterTranslator registers a translator function.
	RegisterTranslator(translator Translator)
	// Load loads the server options from config and custom registered option translators.
	Load(config *config.ClientConfig) ([]client.Option, error)
}

type DefaultOptionLoader struct {
	translators []Translator
}

func NewOptionLoader() OptionLoader {
	return &DefaultOptionLoader{
		translators: make([]Translator, 0),
	}
}

// RegisterTranslator registers a translator function.
// If the translator function has been registered, both will be registered,
// and the translator functions will be called in the order of registration.
func (loader *DefaultOptionLoader) RegisterTranslator(translator Translator) {
	loader.translators = append(loader.translators, translator)
}

func (loader *DefaultOptionLoader) Load(config *config.ClientConfig) ([]client.Option, error) {
	if config == nil {
		return nil, fmt.Errorf("client config not set")
	}
	var translatorsList []Translator
	// basic options
	if config.TransportProtocol != 0 {
		translatorsList = append(translatorsList, translator.TransportProtocolTranslator)
	}
	if config.DestService != "" {
		translatorsList = append(translatorsList, translator.DestServiceTranslator)
	}
	if config.HostPorts != nil {
		translatorsList = append(translatorsList, translator.HostPortsTranslator)
	}
	if config.LongConnection != nil {
		translatorsList = append(translatorsList, translator.LongConnectionTranslator)
	}
	if config.MuxConnection != 0 {
		translatorsList = append(translatorsList, translator.MuxConnectionTranslator)
	}
	if config.RPCTimeout != nil {
		translatorsList = append(translatorsList, translator.RPCTimeoutTranslator)
	}
	if config.ConnectTimeout != nil {
		translatorsList = append(translatorsList, translator.ConnectTimeoutTranslator)
	}
	if config.Tag != nil {
		translatorsList = append(translatorsList, translator.TagTranslator)
	}
	if config.StatsLevel != 0 {
		translatorsList = append(translatorsList, translator.StatsLevelTranslator)
	}
	if config.ConnReporterEnabled {
		translatorsList = append(translatorsList, translator.ConnReporterEnabledTranslator)
	}
	if config.GRPCConnPoolSize != 0 {
		translatorsList = append(translatorsList, translator.GRPCConnPoolSizeTranslator)
	}
	if config.GRPCWriteBufferSize != 0 {
		translatorsList = append(translatorsList, translator.GRPCWriteBufferSizeTranslator)
	}
	if config.GRPCReadBufferSize != 0 {
		translatorsList = append(translatorsList, translator.GRPCReadBufferSizeTranslator)
	}
	if config.GRPCInitialWindowSize != 0 {
		translatorsList = append(translatorsList, translator.GRPCInitialWindowSizeTranslator)
	}
	if config.GRPCInitialConnWindowSize != 0 {
		translatorsList = append(translatorsList, translator.GRPCInitialConnWindowSizeTranslator)
	}
	if config.GRPCMaxHeaderListSize != 0 {
		translatorsList = append(translatorsList, translator.GRPCMaxHeaderListSizeTranslator)
	}
	if config.GRPCKeepaliveParams != nil {
		translatorsList = append(translatorsList, translator.GRPCKeepaliveParamsTranslator)
	}
	// advanced options
	if config.HTTPConnection {
		translatorsList = append(translatorsList, translator.HTTPConnectionTranslator)
	}
	if config.ClientBasicInfo != nil {
		translatorsList = append(translatorsList, translator.ClientBasicInfoTranslator)
	}
	// Add the custom registered option translators behind the default translators.
	loader.translators = append(translatorsList, loader.translators...)

	var options []client.Option
	for _, trans := range loader.translators {
		options = append(options, trans(config))
	}
	return options, nil
}
