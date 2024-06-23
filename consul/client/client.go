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
	MyTranslators                    map[string]Translator
	MyStreamTranslators              map[string]StreamTranslator
	MyCallOptionMapTranslators       map[string]CallOptionMapTranslator
	MyCallOptionTranslators          map[string]CallOptionTranslator
	MyStreamCallOptionMapTranslators map[string]StreamCallOptionMapTranslator
	MyStreamCallOptionTranslators    map[string]StreamCallOptionTranslator
	ShouldResultRetry                *retry.ShouldResultRetry
}

func NewLoader(clientServiceName, serverServiceName string, reader *ConsulReader, opts LoaderOptions) (*ConsulLoader, error) {

	// Register all translators
	translators := map[string]Translator{
		"basicInfo":            basicInfoTranslator,
		"hostPorts":            hostPortsTranslator,
		"destService":          destServiceTranslator,
		"protocol":             protocolTranslator,
		"connection":           connectionTranslator,
		"failureRetry":         failureRetryTranslator,
		"specifiedResultRetry": specifiedResultRetryTranslator,
		"backupRequest":        backupRequestTranslator,
		"rpcTimeout":           rpcTimeoutTranslator,
		"connectionTimeout":    connectionTimeoutTranslator,
		"tags":                 tagsTranslator,
		"statsLevel":           statsLevelTranslator,
		"grpc":                 grpcTranslator,
	}
	streamTranslators := map[string]StreamTranslator{
		"streamBasicInfo":         streamBasicInfoTranslator,
		"streamHostPorts":         streamHostPortsTranslator,
		"streamDestService":       streamDestServiceTranslator,
		"streamConnectionTimeout": streamConnectionTimeoutTranslator,
		"streamTags":              streamTagsTranslator,
		"streamStatsLevel":        streamStatsLevelTranslator,
		"streamGrpc":              streamGrpcTranslator,
	}
	callOptionMapTranslators := map[string]CallOptionMapTranslator{
		"callOptionHostPorts": callOptionHostPortsTranslator,
		"callOptionUrls":      callOptionUrlsTranslator,
		"callOptionTags":      callOptionTagsTranslator,
	}
	callOptionTranslators := map[string]CallOptionTranslator{
		"callOptionRPCTimeout":        callOptionRPCTimeoutTranslator,
		"callOptionConnectionTimeout": callOptionConnectionTimeoutTranslator,
		"callOptionHTTPHostTimeout":   callOptionHTTPHostTimeoutTranslator,
		"callOptionRetryPolicy":       callOptionRetryPolicyTranslator,
		"callOptionGRPCCompressor":    callOptionGRPCCompressorTranslator,
	}
	streamCallOptionMapTranslators := map[string]StreamCallOptionMapTranslator{
		"streamCallOptionHostPorts": streamCallOptionHostPortsTranslator,
		"streamCallOptionUrls":      streamCallOptionUrlsTranslator,
		"streamCallOptionTags":      streamCallOptionTagsTranslator,
	}
	streamCallOptionTranslators := map[string]StreamCallOptionTranslator{
		"streamCallOptionConnectionTimeout": streamCallOptionConnectionTimeoutTranslator,
		"streamCallOptionGRPCCompressor":    streamCallOptionGRPCCompressorTranslator,
	}

	loader := &ConsulLoader{
		translators:                    translators,
		streamTranslators:              streamTranslators,
		callOptionMapTranslators:       callOptionMapTranslators,
		callOptionTranslators:          callOptionTranslators,
		streamCallOptionMapTranslators: streamCallOptionMapTranslators,
		streamCallOptionTranslators:    streamCallOptionTranslators,
		clientServiceName:              clientServiceName,
		serverServiceName:              serverServiceName,
		reader:                         reader,
		shouldResultRetry:              opts.ShouldResultRetry,
	}

	for name, translator := range opts.MyTranslators {
		loader.RegisterTranslator(name, translator)
	}
	for name, translator := range opts.MyStreamTranslators {
		loader.RegisterStreamTranslator(name, translator)
	}
	for name, translator := range opts.MyCallOptionMapTranslators {
		loader.RegisterCallOptionMapTranslator(name, translator)
	}
	for name, translator := range opts.MyCallOptionTranslators {
		loader.RegisterCallOptionTranslator(name, translator)
	}
	for name, translator := range opts.MyStreamCallOptionMapTranslators {
		loader.RegisterStreamCallOptionMapTranslator(name, translator)
	}
	for name, translator := range opts.MyStreamCallOptionTranslators {
		loader.RegisterStreamCallOptionTranslator(name, translator)
	}

	return loader, nil
}
