package server

import (
	kitexserver "github.com/cloudwego/kitex/server"
)

type serverTranslator func(config map[string]interface{}) ([]kitexserver.Option, error)

// interfaces
// type translator func(config map[string]interface{}) (interface{}, error)
type Loader interface {
	SetSource(reader interface{}) error
	Load() error
	GetOptions() (interface{}, error)
	RegisterTranslator(fieldName string, translator serverTranslator) error
	DeregisterTranslator(fieldName string) error
	AddDefaultOptions(opts ...interface{})
}

type serverLoader struct {
	configSources []ymlreader
	options       []kitexserver.Option
	translators   map[string]serverTranslator
}
