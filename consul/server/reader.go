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

package server

import (
	"bytes"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/hashicorp/consul/api"
	"text/template"
	"time"
)

type ConfigType string

const (
	JSON                      ConfigType = "json"
	YAML                      ConfigType = "yaml"
	ConsulDefaultConfigAddr              = "127.0.0.1:8500"
	ConsulDefaultConfigPrefix            = "KitexConfig"
	ConsulDefaultTimeout                 = 5 * time.Second
	ConsulDefaultDataCenter              = "dc1"
	ConsulDefaultServerPath              = "/{{.ServerServiceName}}"
	ConsulDefaultConfigType              = JSON
)

type Reader interface {
	SetDecoder(decoder ConfigParser) error
	ReadToConfig(p *Path) error
	GetConfig() ConsulConfig
}

type ConsulReader struct {
	config             *ConsulConfig //配置文件读出结果
	parser             ConfigParser  //配置文件解码器
	consulClient       *api.Client
	serverPathTemplate *template.Template
	serverPath         string
	prefix             string
	consulTimeout      time.Duration
	configType         ConfigType
}

type Path struct {
	ServerServiceName string
}

func (r *ConsulReader) SetDecoder(decoder ConfigParser) error {
	r.parser = decoder
	return nil
}
func (r *ConsulReader) ReadToConfig(p *Path) error {
	var err error
	r.serverPath, err = r.render(p, r.serverPathTemplate)
	if err != nil {
		return err
	}
	key := r.prefix + r.serverPath
	_, cancel := context.WithTimeout(context.Background(), r.consulTimeout)
	defer cancel()
	kv := r.consulClient.KV()
	data, _, err := kv.Get(key, nil)
	if err != nil {
		klog.Debugf("[consul] key: %s config get value failed", key)
		return err
	}
	err = r.parser.Decode(r.configType, data.Value, r.config)
	if err != nil {
		return err
	}
	return nil
}
func (r *ConsulReader) GetConfig() ConsulConfig { return *r.config }

func (r *ConsulReader) render(p *Path, t *template.Template) (string, error) {
	var tpl bytes.Buffer
	err := t.Execute(&tpl, p)
	if err != nil {
		return "", err
	}
	return tpl.String(), nil
}
