package server

import (
	"context"
	"encoding/json"
	"fmt"
	etcdServer "github.com/Printemps417/optionloader/etcd/server"
	"github.com/Printemps417/optionloader/utils"
	"github.com/cloudwego/kitex-examples/kitex_gen/api"
	"github.com/cloudwego/kitex-examples/kitex_gen/api/echo"
	"github.com/cloudwego/kitex/pkg/klog"
	kitexserver "github.com/cloudwego/kitex/server"
	"log"
)

const (
	serverServiceName = "serverServiceName"
)

var _ api.Echo = &EchoImpl{}

// EchoImpl implements the last service interface defined in the IDL.
type EchoImpl struct{}

// Echo implements the Echo interface.
func (s *EchoImpl) Echo(ctx context.Context, req *api.Request) (resp *api.Response, err error) {
	klog.Info("echo called")
	return &api.Response{Message: req.Message}, nil
}

// 用户可以自定义读取数据的类型，要求通过Decode返回一个字节流
type myConfigParser struct {
}

func (p *myConfigParser) Decode(data []byte, config etcdServer.EtcdConfig) error {
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
func myTranslator(config *etcdServer.EtcdConfig) ([]kitexserver.Option, error) {
	c := config.MyConfig
	var opts []kitexserver.Option
	//具体处理逻辑
	_ = c
	_ = opts
	return opts, nil
}

func main() {
	readerOptions := etcdServer.ReaderOptions{
		ConfigParser: &myConfigParser{},
		MyConfig:     &myConfig{},
	}
	utils.Printpath()
	reader, err := etcdServer.NewReader(readerOptions)
	//reader, err := etcdClient.NewReader(etcdClient.ReaderOptions{})//使用默认值时的
	if err != nil {
		log.Fatal(err)
		return
	}
	myTranslators := []etcdServer.Translator{myTranslator}
	loader, err := etcdServer.NewLoader(serverServiceName, reader, myTranslators...)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = loader.Load()
	if err != nil {
		log.Fatal(err)
		return
	}
	server := echo.NewServer(
		new(EchoImpl),
		kitexserver.WithSuite(loader.GetSuite()),
	)
	if err := server.Run(); err != nil {
		log.Println("server stopped with error:", err)
	} else {
		log.Println("server stopped")
	}
}
