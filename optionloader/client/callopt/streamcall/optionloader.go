package streamcall

import (
	"fmt"
	"github.com/cloudwego/kitex/client/callopt/streamcall"
	"github.com/kitex-contrib/optionloader/config"
	translator "github.com/kitex-contrib/optionloader/translator/client/callopt/streamcall"
)

type Translator func(config *config.StreamCallConfig) streamcall.Option

type OptionLoader interface {
	// RegisterTranslator registers a translator function.
	RegisterTranslator(translator Translator)
	// Load loads the server options from config and custom registered option translators.
	Load(config *config.StreamCallConfig) ([]streamcall.Option, error)
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

func (loader *DefaultOptionLoader) Load(config *config.StreamCallConfig) ([]streamcall.Option, error) {
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

	var options []streamcall.Option
	for _, trans := range loader.translators {
		options = append(options, trans(config))
	}
	return options, nil
}
