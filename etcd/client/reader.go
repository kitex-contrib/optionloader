package client

type Reader interface {
	SetDecoder(decoder parser) error
	ReadToConfig(data []byte) error
	GetConfig(key string) (map[string]interface{}, error)
	ReadEtcdConfig(key Key) error
}

// EtcdReader yml/client/reader.go
type etcdReader struct {
	config  etcdConfig //配置文件读出结果
	decoder parser     //配置文件解码器
}

type Key struct {
	Prefix string
	Path   string
}

func (r *etcdReader) SetDecoder(decoder parser) error {
	r.decoder = decoder
	return nil
}
func (r *etcdReader) ReadToConfig(data []byte) error {
	//todo
	return nil
}
func (r *etcdReader) GetConfig(key string) (map[string]interface{}, error) {
	//todo
	return nil, nil
}

func (r *etcdReader) ReadEtcdConfig(key Key) error {
	//todo
	return nil
}
