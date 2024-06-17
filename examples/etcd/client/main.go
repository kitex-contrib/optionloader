package main

import (
	"context"
	"encoding/json"
	"fmt"
	etcdClient "github.com/Printemps417/optionloader/etcd/client"
	"github.com/Printemps417/optionloader/utils"
	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	kitexclient "github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/klog"
	"log"
	"time"
)

const (
	serverServiceName = "serverServiceName"
	clientServiceName = "clientServiceName"
)

// 用户可以自定义读取数据的类型，要求通过Decode返回一个字节流
type myConfigParser struct {
}

func (p *myConfigParser) Decode(data []byte, config etcdClient.EtcdConfig) error {
	return json.Unmarshal(data, &config)
}

// 用户可以自定义新增Config文件结构，并且默认的的Config文件结构仍然存在
type myConfig struct {
	configOne *string `mapstructure:"configOne"`
	configTwo *string `mapstructure:"configTwo"`
}

func (r *myConfig) String() string {
	return fmt.Sprintf(
		"configOne: %s\n"+
			"configTwo: %s\n", *r.configOne, *r.configTwo)
}

// 用户可自定义Translator，用于将myConfig解析成Options
func myTranslator(config *etcdClient.EtcdConfig) ([]kitexclient.Option, error) {
	c := config.MyConfig
	opts := []kitexclient.Option{}
	//具体处理逻辑
	_ = c
	_ = opts
	return opts, nil
}

func main() {
	readerOptions := etcdClient.ReaderOptions{
		ConfigParser: &myConfigParser{},
		MyConfig:     &myConfig{},
	}
	utils.Printpath()
	reader, err := etcdClient.NewReader(readerOptions)
	//reader, err := etcdClient.NewReader(etcdClient.ReaderOptions{})//使用默认值时的
	if err != nil {
		log.Fatal(err)
		return
	}
	myTranslators := []etcdClient.ClientTranslator{myTranslator}
	loader, err := etcdClient.NewLoader(clientServiceName, serverServiceName, reader, myTranslators...)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = loader.Load()
	if err != nil {
		log.Fatal(err)
		return
	}
	client, err := echo.NewClient(
		serverServiceName,
		kitexclient.WithSuite(loader.GetSuite()),
	)
	if err != nil {
		log.Fatal(err)
	}
	for {
		req := &api.Request{Message: "my request"}
		resp, err := client.Echo(context.Background(), req)
		if err != nil {
			klog.Errorf("take request error: %v", err)
		} else {
			klog.Infof("receive response %v", resp)
		}
		time.Sleep(time.Second * 10)
	}
}
