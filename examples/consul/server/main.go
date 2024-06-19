package main

import (
	"encoding/json"
	"fmt"
	consulServer "github.com/Printemps417/optionloader/consul/server"
	"github.com/Printemps417/optionloader/utils"
	kitexserver "github.com/cloudwego/kitex/server"
	"gopkg.in/yaml.v3"
	"log"
)

const (
	serverServiceName = "echo_server_service"
)

// 用户可以自定义读取数据的类型，要求通过Decode返回一个字节流
type myConfigParser struct {
}

func (p *myConfigParser) Decode(configType consulServer.ConfigType, data []byte, config *consulServer.ConsulConfig) error {
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

func (r *myConfig) String() string {
	var output string

	if r.ConfigOne != nil {
		output += fmt.Sprintf("ConfigOne: %s\n", *r.ConfigOne)
	}

	if r.ConfigTwo != nil {
		output += fmt.Sprintf("ConfigTwo: %v\n", r.ConfigTwo)
	}

	return output
}

// 用户可自定义Translator，用于将myConfig解析成Options
func myTranslator(config *consulServer.ConsulConfig) ([]kitexserver.Option, error) {
	c := config.MyConfig
	if c == nil {
		return nil, nil
	}
	var opts []kitexserver.Option
	//具体处理逻辑
	_ = opts
	fmt.Println("myConfigTranslator run! myConfig:" + c.String())
	return opts, nil
}

func main() {
	readerOptions := consulServer.ReaderOptions{
		ConfigParser: &myConfigParser{},
		MyConfig:     &myConfig{},
	}
	utils.Printpath()
	reader, err := consulServer.NewReader(readerOptions)
	//reader, err := etcdClient.NewReader(etcdClient.ReaderOptions{})//使用默认值时的
	if err != nil {
		log.Fatal(err)
		return
	}
	myTranslators := []consulServer.Translator{myTranslator}
	loader, err := consulServer.NewLoader(serverServiceName, reader, myTranslators...)
	if err != nil {
		log.Fatal(err)
		return
	}
	err = loader.Load()
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println("Options: ", loader.GetSuite().Options())
	config, _ := reader.GetConfig()
	fmt.Print("Config:", config.String())
	//	server := echo.NewServer(
	//		new(EchoImpl),
	//		kitexserver.WithSuite(loader.GetSuite()),
	//	)
	//	if err := server.Run(); err != nil {
	//		log.Println("server stopped with error:", err)
	//	} else {
	//		log.Println("server stopped")
	//	}
}
