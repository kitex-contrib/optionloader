package server

import (
	"github.com/cloudwego/kitex/pkg/klog"
	kitexserver "github.com/cloudwego/kitex/server"
)

type Translator func(config *ConsulConfig) ([]kitexserver.Option, error)

type Loader interface {
	Load() error
	GetSuite() *ConsulServerSuite
}

type ConsulLoader struct {
	reader            *ConsulReader
	options           []kitexserver.Option
	translators       []Translator
	ServerServiceName string
	suite             *ConsulServerSuite
}

func (l *ConsulLoader) Load() error {
	path := Path{ServerServiceName: l.ServerServiceName}
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
	l.suite = &ConsulServerSuite{
		opts: l.options,
	}
	return nil
}

func (l *ConsulLoader) GetSuite() *ConsulServerSuite {
	// 返回当前的 options
	return l.suite
}
