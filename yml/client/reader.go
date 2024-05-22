package client

// decoder可选为用户自定义，以支持多种配置文件
type Decoder func(kind string, data []byte, config interface{}) error

// Reader defines the interface for reading configuration from different sources.

// interfaces
type Reader interface {
	SetDecoder(decoder Decoder) error
	ReadToConfig(data []byte) error
	GetConfig(key string) (map[string]interface{}, error)
}

// yml/client/reader.go
type ymlreader struct {
	config   map[string]interface{} //配置文件读出结果
	filePath string                 //配置文件路径
	decoder  Decoder                //配置文件解码器
}
