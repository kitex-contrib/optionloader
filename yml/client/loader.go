package client

import (
	kitexclient "github.com/cloudwego/kitex/client"
)

type clientTranslator func(FieldConfig interface{}) ([]kitexclient.Option, error)

// interfaces
// type translator func(config map[string]interface{}) (interface{}, error)
type Loader interface {
	SetSource(reader ymlReader) error
	Load() error
	GetOptions() ([]kitexclient.Option, error)
	RegisterTranslator(fieldName string, translator clientTranslator) error
	DeregisterTranslator(fieldName string) error
	AddDefaultOptions(opts ...interface{})
}

type clientLoader struct {
	configSource ymlReader
	options      []kitexclient.Option
	translators  map[string]clientTranslator
}

func (loader *clientLoader) SetSource(reader ymlReader) error {
	loader.configSource = reader
	return nil
}
func (loader *clientLoader) Load() error {
	for field, translator := range loader.translators {
		// 通过字段名获取字段值
		fieldConfig, err := loader.configSource.GetConfigByField(field)
		if err != nil {
			return err
		}
		// 通过字段名获取选项
		opts, err := translator(fieldConfig)
		if err != nil {
			return err
		}
		loader.options = append(loader.options, opts...)
	}
	return nil
}

// 实现 Loader 接口的 GetOptions 方法。
func (l *clientLoader) GetOptions() (interface{}, error) {
	// 返回当前的 options
	return l.options, nil
}

// 实现 Loader 接口的 RegisterTranslator 方法。
func (l *clientLoader) RegisterTranslator(fieldName string, translator clientTranslator) error {
	// 注册字段名到 client 选项的转换器
	if l.translators == nil {
		l.translators = make(map[string]clientTranslator)
	}
	l.translators[fieldName] = translator
	return nil
}

// 实现 Loader 接口的 DeregisterTranslator 方法。
func (l *clientLoader) DeregisterTranslator(fieldName string) error {
	// 注销字段名到 client 选项的转换器
	delete(l.translators, fieldName)
	return nil
}

// 实现 Loader 接口的 AddDefaultOptions 方法。
func (l *clientLoader) AddDefaultOptions(opts []kitexclient.Option) {
	// 添加默认选项
	for _, opt := range opts {
		l.options = append(l.options, opt)
	}
}
func portTranslator(p port) ([]kitexclient.Option, error) {
	res := []kitexclient.Option{}
	s := string(p)
	res = append(res, kitexclient.WithHostPorts(s))
	return res, nil
}
func NewClientLoader() (*clientLoader, error) {
	loader := &clientLoader{
		translators: make(map[string]clientTranslator),
	}

	err := loader.RegisterTranslator("port", portTranslator)
	if err != nil {
		return nil, err
	}

	return loader, nil
}
