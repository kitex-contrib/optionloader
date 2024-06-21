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
	kitexclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
)

type Translator func(config *ConsulConfig) ([]kitexclient.Option, error)

type Loader interface {
	Load() error
	GetSuite() *ConsulClientSuite
}

type ConsulLoader struct {
	reader            *ConsulReader
	options           []kitexclient.Option
	translators       []Translator
	clientServiceName string
	serverServiceName string
	suite             *ConsulClientSuite
	shouldResultRetry *retry.ShouldResultRetry
}

func (l *ConsulLoader) Load() error {
	path := Path{ClientServiceName: l.clientServiceName, ServerServiceName: l.serverServiceName}
	err := l.reader.ReadToConfig(&path)
	if err != nil {
		return err
	}
	config, err := l.reader.GetConfig()
	if err != nil {
		return err
	}
	if l.shouldResultRetry != nil {
		if config.FailureRetry != nil {
			config.FailureRetry.ShouldResultRetry = l.shouldResultRetry
		} else {
			config.ShouldResultRetry = l.shouldResultRetry
		}
	}
	for _, translator := range l.translators {
		opts, err := translator(config)
		if err != nil {
			klog.Errorf(err.Error())
			continue
		}
		l.options = append(l.options, opts...)
	}
	l.suite = &ConsulClientSuite{
		opts: l.options,
	}
	return nil
}

func (l *ConsulLoader) GetSuite() *ConsulClientSuite {
	// 返回当前的 options
	return l.suite
}
