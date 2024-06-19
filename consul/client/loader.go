package client

import (
	kitexclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
)

type Translator func(config *EtcdConfig) ([]kitexclient.Option, error)

type Loader interface {
	Load() error
	GetSuite() EtcdClientSuite
}

type EtcdLoader struct {
	reader            *EtcdReader
	options           []kitexclient.Option
	translators       []Translator
	ClientServiceName string
	ServerServiceName string
	suite             *EtcdClientSuite
}

func (l *EtcdLoader) Load() error {
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
	l.suite = &EtcdClientSuite{
		opts: l.options,
	}
	return nil
}

func (l *EtcdLoader) GetSuite() *EtcdClientSuite {
	// 返回当前的 options
	return l.suite
}
