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
	"fmt"
	ecli "go.etcd.io/etcd/client/v3"
	"text/template"
	"time"
)

const (
	EtcdDefaultNode         = "http://127.0.0.1:2379"
	EtcdDefaultConfigPrefix = "/KitexConfig"
	EtcdDefaultTimeout      = 5 * time.Second
	EtcdServerDefaultPath   = "/{{.ServerServiceName}}"
)

type Reader interface {
	SetDecoder(decoder ConfigParser) error
	ReadToConfig(p *Path) error
	GetConfig() (*EtcdConfig, error)
}

type EtcdReader struct {
	config             *EtcdConfig  //配置文件读出结果
	parser             ConfigParser //配置文件解码器
	etcdClient         *ecli.Client
	serverPathTemplate *template.Template
	serverPath         string
	prefix             string
	etcdTimeout        time.Duration
}

type Path struct {
	ServerServiceName string
}

func (r *EtcdReader) SetDecoder(decoder ConfigParser) error {
	r.parser = decoder
	return nil
}
func (r *EtcdReader) ReadToConfig(p *Path) error {
	var err error
	r.serverPath, err = r.render(p, r.serverPathTemplate)
	if err != nil {
		return err
	}
	key := r.prefix + r.serverPath
	ctx2, cancel := context.WithTimeout(context.Background(), r.etcdTimeout)
	defer cancel()
	data, err := r.etcdClient.Get(ctx2, key)
	if err != nil {
		fmt.Printf("[etcd] key: %s config get value failed", key)
		return err
	}
	err = r.parser.Decode(data.Kvs[0].Value, r.config)
	if err != nil {
		return err
	}
	return nil
}
func (r *EtcdReader) GetConfig() (*EtcdConfig, error) {
	return r.config, nil
}

func (r *EtcdReader) render(p *Path, t *template.Template) (string, error) {
	var tpl bytes.Buffer
	err := t.Execute(&tpl, p)
	if err != nil {
		return "", err
	}
	return tpl.String(), nil
}
