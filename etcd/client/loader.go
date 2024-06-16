package client

import (
	kitexclient "github.com/cloudwego/kitex/client"
)

type ClientTranslator func(config *EtcdConfig) ([]kitexclient.Option, error)

// interfaces
// type translator func(config map[string]interface{}) (interface{}, error)
type Loader interface {
	Load() error
	GetOptions() ([]kitexclient.Option, error)
}

type EtcdLoader struct {
	reader            *EtcdReader
	options           []kitexclient.Option
	translators       []ClientTranslator
	ClientServiceName string
	ServerServiceName string
}

func (l *EtcdLoader) Load() error {
	config, err := l.reader.GetConfig()
	if err != nil {
		return err
	}
	for _, translator := range l.translators {
		// 通过字段名获取选项
		opts, err := translator(config)
		if err != nil {
			continue
		}
		l.options = append(l.options, opts...)
	}
	return nil
}

// 实现 Loader 接口的 GetOptions 方法。
func (l *EtcdLoader) GetOptions() ([]kitexclient.Option, error) {
	// 返回当前的 options
	return l.options, nil
}
