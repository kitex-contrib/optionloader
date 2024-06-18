# .github
example中为使用示例，完整example代码地址：https://github.com/zhu-mi-shan/optionloader_example
使用该组件需要安装etcd。
etcd中存储的配置文件例子如下：
server:
```
json_data='{"ServerBasicInfo": {"ServiceName": "echo_server_service","Method": "method1","Tags": {"tag1": "v1","tag2": "v2"}},"ServiceAddr": [{"Network": "tcd","Address": "127.0.0.1:8889"}],"MuxTransport": true,"MyConfig":{"configOne": "This is configOne","configTwo": ["Welcome","to","configTwo","!"]}}'

etcdctl put "/KitexConfig/echo_server_service" "$(echo -n $json_data)"
```
client:
```
json_data='{"ClientBasicInfo":{"ServiceName": "echo_client_service","Method": "method1","Tags": {"tag1": "v1","tag2": "v2"}},"HostPorts": ["0.0.0.0:8888","0.0.0.0:8889"],"DestService": "echo_server_service","Protocol": "HTTP","Connection":{"Method": "LongConnection","LongConnection":{"MinIdlePerAddress": 1,"MaxIdlePerAddress": 10,"MaxIdleGlobal": 100,"MaxIdleTimeout": "1m"},"MuxConnection":{"ConnNum": 3}},"MyConfig":{"configOne": "This is configOne","configTwo": ["Welcome","to","configTwo","!"]}}'

etcdctl put "/KitexConfig/echo_client_service/echo_server_service" "$(echo -n $json_data)"
```
服务端，客户端均支持自定义新的配置文件结构config，自定义新的数据读取方式decoder, 以及新的解析option方法translator
server端示例：
```
const (
	serverServiceName = "echo_server_service"
)

// 用户可以自定义读取数据的类型，要求通过Decode返回一个字节流
type myConfigParser struct {
}

func (p *myConfigParser) Decode(data []byte, config *etcdServer.EtcdConfig) error {
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
func myTranslator(config *etcdServer.EtcdConfig) ([]kitexserver.Option, error) {
	c := config.MyConfig
	var opts []kitexserver.Option
	//具体处理逻辑
	_ = opts
	fmt.Println("myConfigTranslator run! myConfig:" + c.String())
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
	fmt.Println("Options: ", loader.GetSuite().Options())
	config, _ := reader.GetConfig()
	fmt.Print("Config:", config.String())
	svr := example.NewServer(new(TestServiceImpl), kitexserver.WithSuite(loader.GetSuite()))

	err = svr.Run()

	if err != nil {
		log.Println(err.Error())
	}
}
```
client端示例：
```
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
	if c == nil {
		return nil, nil
	}
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

	c, err := example.NewClient("echo_server_service", kitexclient.WithSuite(loader.GetSuite()))
	if err != nil {
		log.Fatal(err)
	}
	req := examplegen.Req{
		Id: 123,
	}
	resp, err := c.Test(context.Background(), &req)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(resp)
}
```
