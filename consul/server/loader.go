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
	"github.com/cloudwego/kitex/pkg/klog"
	kitexserver "github.com/cloudwego/kitex/server"
)

type Loader interface {
	Load() error
	GetSuite() ConsulServerSuite
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
	config := l.reader.GetConfig()
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

func (l *ConsulLoader) GetSuite() ConsulServerSuite {
	// 返回当前的 options
	return *l.suite
}
