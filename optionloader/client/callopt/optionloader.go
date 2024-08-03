package callopt

import (
	"fmt"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/kitex-contrib/optionloader/config"
	translator "github.com/kitex-contrib/optionloader/translator/client/callopt"
)

type Translator func(config *config.CalloptConfig) callopt.Option

type OptionLoader interface {
	// RegisterTranslator registers a translator function.
	RegisterTranslator(translator Translator)
	// Load loads the server options from config and custom registered option translators.
	Load(config *config.CalloptConfig) ([]callopt.Option, error)
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

func (loader *DefaultOptionLoader) Load(config *config.CalloptConfig) ([]callopt.Option, error) {
	if config == nil {
		return nil, fmt.Errorf("client config not set")
	}
	var translatorsList []Translator

	if config.HostPort != "" {
		translatorsList = append(translatorsList, translator.HostPortTranslator)
	}
	if config.URL != "" {
		translatorsList = append(translatorsList, translator.URLTranslator)
	}
	if config.HTTPHost != "" {
		translatorsList = append(translatorsList, translator.HTTPHostTranslator)
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
	if config.GRPCCompressor != "" {
		translatorsList = append(translatorsList, translator.GRPCCompressorTranslator)
	}

	// Add the custom registered option translators behind the default translators.
	loader.translators = append(translatorsList, loader.translators...)

	var options []callopt.Option
	for _, trans := range loader.translators {
		options = append(options, trans(config))
	}
	return options, nil
}
