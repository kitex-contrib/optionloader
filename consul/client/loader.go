package client

import (
	kitexclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
)

type Translator func(config *ConsulConfig) ([]kitexclient.Option, error)

type Loader interface {
	Load() error
	GetSuite() *ConsulClientSuite
}

type ConsulLoader struct {
	reader            *ConsulReader
	options           []kitexclient.Option
	translators       []Translator
	ClientServiceName string
	ServerServiceName string
	suite             *ConsulClientSuite
}

func (l *ConsulLoader) Load() error {
	path := Path{ClientServiceName: l.ClientServiceName, ServerServiceName: l.ServerServiceName}
	err := l.reader.ReadToConfig(&path)
	if err != nil {
		return err
	}
	config, err := l.reader.GetConfig()
	if err != nil {
		return err
	}
	for _, translator := range l.translators {
		opts, err := translator(config)
		if err != nil {
			klog.Errorf(err.Error())
			continue
		}
		l.options = append(l.options, opts...)
	}
	l.suite = &ConsulClientSuite{
		opts: l.options,
	}
	return nil
}

func (l *ConsulLoader) GetSuite() *ConsulClientSuite {
	// 返回当前的 options
	return l.suite
}
