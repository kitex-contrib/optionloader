package client

import (
	kitexclient "github.com/cloudwego/kitex/client"
)

type clientTranslator func(config map[string]interface{}) ([]kitexclient.Option, error)

// interfaces
// type translator func(config map[string]interface{}) (interface{}, error)
type Loader interface {
	SetSource(reader interface{}) error
	Load() error
	GetOptions() (interface{}, error)
	RegisterTranslator(fieldName string, translator clientTranslator) error
	DeregisterTranslator(fieldName string) error
	AddDefaultOptions(opts ...interface{})
}

type clientLoader struct {
	configSources []ymlreader
	options       []kitexclient.Option
	translators   map[string]clientTranslator
}
