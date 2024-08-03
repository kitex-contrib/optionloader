package server

import (
	"fmt"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/optionloader/config"
	translator "github.com/kitex-contrib/optionloader/translator/server"
)

type Translator func(config *config.ServerConfig) server.Option

type OptionLoader interface {
	// RegisterTranslator registers a translator function.
	RegisterTranslator(translator Translator)
	// Load loads the server options from config and custom registered option translators.
	Load(config *config.ServerConfig) ([]server.Option, error)
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

// Load loads the server options from config and custom registered option translators.
// The custom registered option translators will have higher priority than the option translators from the config.
func (loader *DefaultOptionLoader) Load(config *config.ServerConfig) ([]server.Option, error) {
	if config == nil {
		return nil, fmt.Errorf("server config not set")
	}
	var translatorsList []Translator
	// basic options
	if config.MuxTransport {
		translatorsList = append(translatorsList, translator.MuxTransportTranslator)
	}
	if config.ReadWriteTimeout != nil {
		translatorsList = append(translatorsList, translator.ReadWriteTimeoutTranslator)
	}
	if config.ExitWaitTime != nil {
		translatorsList = append(translatorsList, translator.ExitWaitTimeTranslator)
	}
	if config.MaxConnIdleTime != nil {
		translatorsList = append(translatorsList, translator.MaxConnIdleTimeTranslator)
	}
	if config.StatsLevel != 0 {
		translatorsList = append(translatorsList, translator.StatsLevelTranslator)
	}
	if config.ServiceAddr != nil {
		translatorsList = append(translatorsList, translator.ServiceAddrTranslator)
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
	if config.GRPCKeepaliveParams != nil {
		translatorsList = append(translatorsList, translator.GRPCKeepaliveParamsTranslator)
	}
	if config.GRPCKeepaliveEnforcementPolicy != nil {
		translatorsList = append(translatorsList, translator.GRPCKeepaliveEnforcementPolicyTranslator)
	}
	if config.GRPCMaxConcurrentStreams != 0 {
		translatorsList = append(translatorsList, translator.GRPCMaxConcurrentStreamsTranslator)
	}
	if config.GRPCMaxHeaderListSize != 0 {
		translatorsList = append(translatorsList, translator.GRPCMaxHeaderListSizeTranslator)
	}
	if config.ContextBackup != nil {
		translatorsList = append(translatorsList, translator.ContextBackupTranslator)
	}
	if config.RefuseTrafficWithoutServiceName {
		translatorsList = append(translatorsList, translator.RefuseTrafficWithoutServiceNameTranslator)
	}
	if config.EnableContextTimeout {
		translatorsList = append(translatorsList, translator.EnableContextTimeoutTranslator)
	}
	// advanced options
	if config.ServerBasicInfo != nil {
		translatorsList = append(translatorsList, translator.ServerBasicInfoTranslator)
	}
	if config.ReusePort {
		translatorsList = append(translatorsList, translator.ReusePortTranslator)
	}
	// stream options
	if config.CompatibleMiddlewareForUnary {
		translatorsList = append(translatorsList, translator.CompatibleMiddlewareForUnaryTranslator)
	}
	// Add the custom registered option translators behind the default translators.
	loader.translators = append(translatorsList, loader.translators...)

	var options []server.Option
	for _, trans := range loader.translators {
		options = append(options, trans(config))
	}
	return options, nil
}
