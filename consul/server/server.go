package server

import (
	"github.com/hashicorp/consul/api"
	"go.uber.org/zap"
	"text/template"
	"time"
)

// Options etcd config options. All the fields have default value.
type ReaderOptions struct {
	Addr         string
	Prefix       string
	PathFormat   string
	DataCenter   string
	TimeOut      time.Duration
	NamespaceId  string
	Token        string
	Partition    string
	LoggerConfig *zap.Config
	ConfigParser ConfigParser
	ConfigType   ConfigType
	MyConfig     Config
}

func NewReader(opts ReaderOptions) (*ConsulReader, error) {
	if opts.Addr == "" {
		opts.Addr = ConsulDefaultConfigAddr
	}
	if opts.Prefix == "" {
		opts.Prefix = ConsulDefaultConfigPrefix
	}
	if opts.ConfigParser == nil {
		opts.ConfigParser = &defaultParser{}
	}
	if opts.TimeOut == 0 {
		opts.TimeOut = ConsulDefaultTimeout
	}
	if opts.PathFormat == "" {
		opts.PathFormat = ConsulDefaultServerPath
	}
	if opts.DataCenter == "" {
		opts.DataCenter = ConsulDefaultDataCenter
	}
	if opts.ConfigType == "" {
		opts.ConfigType = ConsulDefaultConfigType
	}
	consulClient, err := api.NewClient(&api.Config{
		Address:    opts.Addr,
		Datacenter: opts.DataCenter,
		Token:      opts.Token,
		Namespace:  opts.NamespaceId,
		Partition:  opts.Partition,
	})
	if err != nil {
		return nil, err
	}
	serverPathTemplate, err := template.New("serverName").Parse(opts.PathFormat)
	if err != nil {
		return nil, err
	}
	r := &ConsulReader{
		config:             &ConsulConfig{MyConfig: opts.MyConfig}, //配置文件读出结果
		parser:             opts.ConfigParser,                      //配置文件解码器
		consulClient:       consulClient,
		prefix:             opts.Prefix,
		serverPathTemplate: serverPathTemplate,
		consulTimeout:      opts.TimeOut,
		configType:         opts.ConfigType,
	}

	return r, nil
}

func NewLoader(serverServiceName string, reader *ConsulReader, myTranslators ...Translator) (*ConsulLoader, error) {

	// Register all translators
	translators := []Translator{
		basicInfoTranslator,
		serviceAddrTranslator,
		muxTransportTranslator,
	}

	if len(myTranslators) != 0 {
		translators = append(translators, myTranslators...)
	}

	loader := &ConsulLoader{
		translators:       translators,
		ServerServiceName: serverServiceName,
		reader:            reader,
	}

	return loader, nil
}
