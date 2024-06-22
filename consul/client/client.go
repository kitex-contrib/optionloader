// Copyright 2024 CloudWeGo Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package client

import (
	"github.com/cloudwego/kitex/pkg/retry"
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
		opts.PathFormat = ConsulDefaultClientPath
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
	clientPathTemplate, err := template.New("clientName").Parse(opts.PathFormat)
	if err != nil {
		return nil, err
	}
	r := &ConsulReader{
		config:             &ConsulConfig{MyConfig: opts.MyConfig}, //配置文件读出结果
		parser:             opts.ConfigParser,                      //配置文件解码器
		consulClient:       consulClient,
		prefix:             opts.Prefix,
		clientPathTemplate: clientPathTemplate,
		consulTimeout:      opts.TimeOut,
		configType:         opts.ConfigType,
	}

	return r, nil
}

type LoaderOptions struct {
	MyTranslators     []Translator
	ShouldResultRetry *retry.ShouldResultRetry
}

func NewLoader(clientServiceName, serverServiceName string, reader *ConsulReader, opts LoaderOptions) (*ConsulLoader, error) {

	// Register all translators
	translators := []Translator{
		basicInfoTranslator,
		hostPortsTranslator,
		destServiceTranslator,
		protocolTranslator,
		connectionTranslator,
		failureRetryTranslator,
		specifiedResultRetryTranslator,
		backupRequestTranslator,
		rpcTimeoutTranslator,
		connectionTimeoutTranslator,
		tagsTranslator,
		statsLevelTranslator,
		grpcConnPoolSizeTranslator,
		grpcWriteBufferSizeTranslator,
	}

	if len(opts.MyTranslators) != 0 {
		translators = append(translators, opts.MyTranslators...)
	}

	loader := &ConsulLoader{
		translators:       translators,
		clientServiceName: clientServiceName,
		serverServiceName: serverServiceName,
		reader:            reader,
		shouldResultRetry: opts.ShouldResultRetry,
	}

	return loader, nil
}
