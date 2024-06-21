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
	"fmt"
	kitexserver "github.com/cloudwego/kitex/server"
)

type Translator func(config *EtcdConfig) ([]kitexserver.Option, error)

type Loader interface {
	Load() error
	GetSuite() *EtcdServerSuite
}

type EtcdLoader struct {
	reader            *EtcdReader
	options           []kitexserver.Option
	translators       []Translator
	ServerServiceName string
	suite             *EtcdServerSuite
}

func (l *EtcdLoader) Load() error {
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
			fmt.Println(err)
			continue
		}
		l.options = append(l.options, opts...)
	}
	l.suite = &EtcdServerSuite{
		opts: l.options,
	}
	return nil
}

func (l *EtcdLoader) GetSuite() *EtcdServerSuite {
	// 返回当前的 options
	return l.suite
}
