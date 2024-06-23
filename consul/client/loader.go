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
	"github.com/cloudwego/kitex/client/streamclient"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/retry"
)

type Loader interface {
	Load() error
	GetCallOpt(name string) interface{}
	GetSuite() ConsulClientSuite
	GetStreamSuite() ConsulClientStreamSuite
	RegisterTranslator(name string, translator Translator)
	RegisterStreamTranslator(name string, streamTranslator StreamTranslator)
	RegisterCallOptionMapTranslator(name string, callOptionMapTranslator CallOptionMapTranslator)
	RegisterCallOptionTranslator(name string, callOptionTranslator CallOptionTranslator)
	RegisterStreamCallOptionMapTranslator(name string, streamCallOptionMapTranslator StreamCallOptionMapTranslator)
	RegisterStreamCallOptionTranslator(name string, streamCallOptionTranslator StreamCallOptionTranslator)
}

type ConsulLoader struct {
	reader                         *ConsulReader
	options                        []kitexclient.Option
	streamOptions                  []streamclient.Option
	callOptions                    map[string]interface{}
	translators                    map[string]Translator
	streamTranslators              map[string]StreamTranslator
	callOptionMapTranslators       map[string]CallOptionMapTranslator
	callOptionTranslators          map[string]CallOptionTranslator
	streamCallOptionMapTranslators map[string]StreamCallOptionMapTranslator
	streamCallOptionTranslators    map[string]StreamCallOptionTranslator
	clientServiceName              string
	serverServiceName              string
	suite                          *ConsulClientSuite
	streamSuite                    *ConsulClientStreamSuite
	shouldResultRetry              *retry.ShouldResultRetry
}

func (l *ConsulLoader) Load() error {
	path := Path{ClientServiceName: l.clientServiceName, ServerServiceName: l.serverServiceName}
	err := l.reader.ReadToConfig(&path)
	if err != nil {
		return err
	}
	config := l.reader.GetConfig()
	if l.shouldResultRetry != nil {
		if config.FailureRetry != nil {
			config.FailureRetry.ShouldResultRetry = l.shouldResultRetry
		} else {
			config.ShouldResultRetry = l.shouldResultRetry
		}
	}
	l.options = []kitexclient.Option{}
	for _, translator := range l.translators {
		opts, err := translator(config)
		if err != nil {
			klog.Errorf(err.Error())
			continue
		}
		l.options = append(l.options, opts...)
	}
	l.streamOptions = []streamclient.Option{}
	for _, translator := range l.streamTranslators {
		opts, err := translator(config)
		if err != nil {
			klog.Errorf(err.Error())
			continue
		}
		l.streamOptions = append(l.streamOptions, opts...)
	}
	l.callOptions = map[string]interface{}{}
	for name, translator := range l.callOptionMapTranslators {
		opts := *translator(config)
		l.callOptions[name] = opts
	}
	for name, translator := range l.callOptionTranslators {
		opts := *translator(config)
		l.callOptions[name] = opts
	}
	for name, translator := range l.streamCallOptionMapTranslators {
		opts := *translator(config)
		l.callOptions[name] = opts
	}
	for name, translator := range l.streamCallOptionTranslators {
		opts := *translator(config)
		l.callOptions[name] = opts
	}
	l.suite = &ConsulClientSuite{
		opts: l.options,
	}
	l.streamSuite = &ConsulClientStreamSuite{
		opts: l.streamOptions,
	}
	return nil
}

func (l *ConsulLoader) GetCallOpt(name string) interface{} {
	return l.callOptions[name]
}

func (l *ConsulLoader) GetSuite() ConsulClientSuite {
	return *l.suite
}

func (l *ConsulLoader) GetStreamSuite() ConsulClientStreamSuite {
	return *l.streamSuite
}

func (l *ConsulLoader) RegisterTranslator(name string, translator Translator) {
	l.translators[name] = translator
}

func (l *ConsulLoader) RegisterStreamTranslator(name string, streamTranslator StreamTranslator) {
	l.streamTranslators[name] = streamTranslator
}

func (l *ConsulLoader) RegisterCallOptionMapTranslator(name string, callOptionMapTranslator CallOptionMapTranslator) {
	l.callOptionMapTranslators[name] = callOptionMapTranslator
}

func (l *ConsulLoader) RegisterCallOptionTranslator(name string, callOptionTranslator CallOptionTranslator) {
	l.callOptionTranslators[name] = callOptionTranslator
}

func (l *ConsulLoader) RegisterStreamCallOptionMapTranslator(name string, streamCallOptionMapTranslator StreamCallOptionMapTranslator) {
	l.streamCallOptionMapTranslators[name] = streamCallOptionMapTranslator
}

func (l *ConsulLoader) RegisterStreamCallOptionTranslator(name string, streamCallOptionTranslator StreamCallOptionTranslator) {
	l.streamCallOptionTranslators[name] = streamCallOptionTranslator
}
