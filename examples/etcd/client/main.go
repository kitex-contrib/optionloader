package main

import (
	"encoding/json"
	"fmt"
	etcdClient "github.com/Printemps417/optionloader/etcd/client"
	"github.com/Printemps417/optionloader/utils"
	kitexclient "github.com/cloudwego/kitex/client"
	"log"
)

const (
	serverServiceName = "echo_server_service"
	clientServiceName = "echo_client_service"
)

// 用户可以自定义读取数据的类型，要求通过Decode返回一个字节流
type myConfigParser struct {
}

func (p *myConfigParser) Decode(data []byte, config *etcdClient.EtcdConfig) error {
	return json.Unmarshal(data, config)
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
func myTranslator(config *etcdClient.EtcdConfig) ([]kitexclient.Option, error) {
	c := config.MyConfig
	opts := []kitexclient.Option{}
	//具体处理逻辑
	_ = opts
	fmt.Println("myConfigTranslator run! myConfig:" + c.String())
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
	myTranslators := []etcdClient.Translator{myTranslator}
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
	fmt.Println("Options: ", loader.GetSuite().Options())
	config, _ := reader.GetConfig()
	fmt.Print("Config:", config.String())
	//client, err := echo.NewClient(
	//	serverServiceName,
	//	kitexclient.WithSuite(loader.GetSuite()),
	//)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//for {
	//	req := &api.Request{Message: "my request"}
	//	resp, err := client.Echo(context.Background(), req)
	//	if err != nil {
	//		klog.Errorf("take request error: %v", err)
	//	} else {
	//		klog.Infof("receive response %v", resp)
	//	}
	//	time.Sleep(time.Second * 10)
	//}
}
