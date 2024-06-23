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

package main

import (
	"encoding/json"
	"fmt"
	kitexclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/client/callopt"
	"github.com/cloudwego/kitex/client/callopt/streamcall"
	"github.com/cloudwego/kitex/client/streamclient"
	consulClient "github.com/kitex-contrib/optionloader/consul/client"
	"github.com/kitex-contrib/optionloader/utils"
	"gopkg.in/yaml.v3"
	"log"
)

const (
	serverServiceName = "echo_server_service"
	clientServiceName = "echo_client_service"
)

// 用户可以自定义读取数据的类型，要求通过Decode返回一个字节流
type myConfigParser struct {
}

func (p *myConfigParser) Decode(configType consulClient.ConfigType, data []byte, config *consulClient.ConsulConfig) error {
	switch configType {
	case "json":
		return json.Unmarshal(data, config)
	case "yaml":
		return yaml.Unmarshal(data, config)
	default:
		return fmt.Errorf("unsupported config data type %s", configType)
	}
}

// 用户可以自定义新增Config文件结构，并且默认的的Config文件结构仍然存在
type myConfig struct {
	ConfigOne *string  `mapstructure:"configOne"`
	ConfigTwo []string `mapstructure:"configTwo"`
}

// 用户可自定义Translator，用于将myConfig解析成Options
func myTranslator(config consulClient.ConsulConfig) ([]kitexclient.Option, error) {
	c := config.MyConfig
	if c == nil {
		return nil, nil
	}
	var opts []kitexclient.Option
	//具体处理逻辑
	_ = opts
	fmt.Println("myConfigTranslator run!")
	return opts, nil
}

func myStreamTranslator(config consulClient.ConsulConfig) ([]streamclient.Option, error) {
	c := config.MyConfig
	if c == nil {
		return nil, nil
	}
	var opts []streamclient.Option
	//具体处理逻辑
	_ = opts
	fmt.Println("myStreamTranslators run!")
	return opts, nil
}

func myCallOptionMapTranslator(config consulClient.ConsulConfig) *map[string]callopt.Option {
	c := config.MyConfig
	if c == nil {
		return nil
	}
	var res map[string]callopt.Option
	fmt.Println("myCallOptionMapTranslators run!")
	return &res
}

func myCallOptionTranslator(config consulClient.ConsulConfig) *callopt.Option {
	c := config.MyConfig
	if c == nil {
		return nil
	}
	var res callopt.Option
	fmt.Println("myCallOptionTranslators run!")
	return &res
}

func myStreamCallOptionMapTranslator(config consulClient.ConsulConfig) *map[string]streamcall.Option {
	c := config.MyConfig
	if c == nil {
		return nil
	}
	var res map[string]streamcall.Option
	fmt.Println("myStreamCallOptionMapTranslators run!")
	return &res
}

func myStreamCallOptionTranslator(config consulClient.ConsulConfig) *streamcall.Option {
	c := config.MyConfig
	if c == nil {
		return nil
	}
	var res streamcall.Option
	fmt.Println("myStreamCallOptionTranslators run!")
	return &res
}

func main() {
	readerOptions := consulClient.ReaderOptions{
		ConfigParser: &myConfigParser{},
		MyConfig:     &myConfig{},
	}
	utils.Printpath()
	reader, err := consulClient.NewReader(readerOptions)
	//reader, err := etcdClient.NewReader(etcdClient.ReaderOptions{})//使用默认值时的
	if err != nil {
		log.Fatal(err)
		return
	}
	loaderOptions := consulClient.LoaderOptions{
		MyTranslators:                    map[string]consulClient.Translator{"name1": myTranslator},
		MyStreamTranslators:              map[string]consulClient.StreamTranslator{"name2": myStreamTranslator},
		MyCallOptionMapTranslators:       map[string]consulClient.CallOptionMapTranslator{"name3": myCallOptionMapTranslator},
		MyCallOptionTranslators:          map[string]consulClient.CallOptionTranslator{"name4": myCallOptionTranslator},
		MyStreamCallOptionMapTranslators: map[string]consulClient.StreamCallOptionMapTranslator{"name5": myStreamCallOptionMapTranslator},
		MyStreamCallOptionTranslators:    map[string]consulClient.StreamCallOptionTranslator{"name6": myStreamCallOptionTranslator},
	}
	loader, err := consulClient.NewLoader(clientServiceName, serverServiceName, reader, loaderOptions)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = loader.Load()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("CallOption: ", loader.GetCallOpt("RPCTimeout"))
	config := reader.GetConfig()
	fmt.Print("Config: ", config.String())
	suite := loader.GetSuite()
	fmt.Print("Suite: ", suite)
	streamSuite := loader.GetStreamSuite()
	fmt.Print("Stream Suite: ", streamSuite)
	opts := suite.Options()
	fmt.Println("Options: ", opts)
	streamOpts := streamSuite.Options()
	fmt.Println("Stream Options: ", streamOpts)
}
