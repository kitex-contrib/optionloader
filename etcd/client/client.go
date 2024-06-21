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
	clientv3 "go.etcd.io/etcd/client/v3"
	"text/template"
	"time"
)

// Options etcd config options. All the fields have default value.
type ReaderOptions struct {
	Node         []string
	Prefix       string
	PathFormat   string
	Timeout      time.Duration
	ConfigParser ConfigParser
	MyConfig     Config
}

func NewReader(opts ReaderOptions) (*EtcdReader, error) {
	if opts.Node == nil {
		opts.Node = []string{EtcdDefaultNode}
	}
	if opts.ConfigParser == nil {
		opts.ConfigParser = &defaultParser{}
	}
	if opts.Prefix == "" {
		opts.Prefix = EtcdDefaultConfigPrefix
	}
	if opts.Timeout == 0 {
		opts.Timeout = EtcdDefaultTimeout
	}
	if opts.PathFormat == "" {
		opts.PathFormat = EtcdClientDefaultPath
	}
	etcdClient, err := clientv3.New(clientv3.Config{
		Endpoints:   opts.Node,
		DialTimeout: opts.Timeout,
	})
	if err != nil {
		return nil, err
	}
	clientPathTemplate, err := template.New("clientName").Parse(opts.PathFormat)
	if err != nil {
		return nil, err
	}
	r := &EtcdReader{
		config:             &EtcdConfig{MyConfig: opts.MyConfig}, //配置文件读出结果
		parser:             opts.ConfigParser,                    //配置文件解码器
		etcdClient:         etcdClient,
		prefix:             opts.Prefix,
		clientPathTemplate: clientPathTemplate,
		etcdTimeout:        opts.Timeout,
	}

	return r, nil
}

func NewLoader(clientServiceName, serverServiceName string, reader *EtcdReader, myTranslators ...Translator) (*EtcdLoader, error) {

	// Register all translators
	translators := []Translator{
		basicInfoTranslator,
		hostPortsTranslator,
		destServiceTranslator,
		protocolTranslator,
		connectionTranslator,
	}

	if len(myTranslators) != 0 {
		translators = append(translators, myTranslators...)
	}

	loader := &EtcdLoader{
		translators:       translators,
		ClientServiceName: clientServiceName,
		ServerServiceName: serverServiceName,
		reader:            reader,
	}

	return loader, nil
}
