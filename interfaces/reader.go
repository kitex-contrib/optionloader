package interfaces

// decoder可选为用户自定义，以支持多种配置文件
type Decoder func(kind string, data []byte, config interface{}) error

// Reader defines the interface for reading configuration from different sources.
type Reader interface {
	SetDecoder(decoder Decoder) error
	ReadToConfig(data []byte) error
	GetConfig(key string) (map[string]interface{}, error)
}
