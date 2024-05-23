package client

import (
	"errors"
	"gopkg.in/yaml.v2"
)

// decoder可选为用户自定义，以支持多种配置文件
type Decoder func(data []byte, config interface{}) error

// Reader defines the interface for reading configuration from different sources.

// interfaces
type Reader interface {
	SetDecoder(decoder Decoder) error
	ReadToConfig(data []byte) error
	GetConfig(key string) (map[string]interface{}, error)
}

// yml/client/reader.go
type ymlReader struct {
	config  YMLConfig //配置文件读出结果
	decoder Decoder   //配置文件解码器
}

func (y *ymlReader) SetDecoder(decoder Decoder) error {
	if decoder == nil {
		return errors.New("decoder cannot be nil")
	}
	y.decoder = decoder
	return nil
}

func (y *ymlReader) ReadToConfig(data []byte) error {
	if y.decoder == nil {
		return errors.New("decoder is not set")
	}
	err := y.decoder(data, &y.config)
	if err != nil {
		return err
	}
	return nil
}

func (y *ymlReader) GetConfig() (YMLConfig, error) {
	return y.config, nil
}

// Default decoder
func defaultDecoder(data []byte, out interface{}) error {
	return yaml.Unmarshal(data, out)
}
