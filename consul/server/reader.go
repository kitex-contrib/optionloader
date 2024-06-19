package server

import (
	"bytes"
	"context"
	"github.com/cloudwego/kitex/pkg/klog"
	ecli "go.etcd.io/etcd/client/v3"
	"text/template"
	"time"
)

const (
	EtcdDefaultNode         = "http://127.0.0.1:2379"
	EtcdDefaultConfigPrefix = "/KitexConfig"
	EtcdDefaultTimeout      = 5 * time.Second
	EtcdServerDefaultPath   = "/{{.ServerServiceName}}"
)

type Reader interface {
	SetDecoder(decoder ConfigParser) error
	ReadToConfig(p *Path) error
	GetConfig() (*EtcdConfig, error)
}

type EtcdReader struct {
	config             *EtcdConfig  //配置文件读出结果
	parser             ConfigParser //配置文件解码器
	etcdClient         *ecli.Client
	clientPathTemplate *template.Template
	clientPath         string
	prefix             string
	etcdTimeout        time.Duration
}

type Path struct {
	ServerServiceName string
}

func (r *EtcdReader) SetDecoder(decoder ConfigParser) error {
	r.parser = decoder
	return nil
}
func (r *EtcdReader) ReadToConfig(p *Path) error {
	var err error
	r.clientPath, err = r.render(p, r.clientPathTemplate)
	if err != nil {
		return err
	}
	key := r.prefix + r.clientPath
	ctx2, cancel := context.WithTimeout(context.Background(), r.etcdTimeout)
	defer cancel()
	data, err := r.etcdClient.Get(ctx2, key)
	if err != nil {
		klog.Debugf("[etcd] key: %s config get value failed", key)
		return err
	}
	err = r.parser.Decode(data.Kvs[0].Value, r.config)
	if err != nil {
		return err
	}
	return nil
}
func (r *EtcdReader) GetConfig() (*EtcdConfig, error) {
	return r.config, nil
}

func (r *EtcdReader) render(p *Path, t *template.Template) (string, error) {
	var tpl bytes.Buffer
	err := t.Execute(&tpl, p)
	if err != nil {
		return "", err
	}
	return tpl.String(), nil
}
