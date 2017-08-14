package etcd

import (
	"log"
	"time"

	"github.com/coreos/etcd/client"
	"net/http"
	"os"
	"github.com/coreos/etcd/pkg/transport"
	"context"
	"auto-ng/config"
)

var globalEndpoints []string
var globalCaFile string
var globalCertFile string
var globalKeyFile string

var globalKapi client.KeysAPI

/**
	初始化 etcd客户端的一些链接设置等等
 */

func init() {
	cfg := config.GetETCDconfig()
	globalEndpoints = cfg.EndPoints
	globalCaFile = cfg.CaFile
	globalCertFile = cfg.CertFile
	globalKeyFile = cfg.KeyFile
}

/**
	获取client，没有创建有就使用
 */

func getKapi() client.KeysAPI {
	if globalKapi == nil {
		transport, err := getTransport()
		if err != nil {
			log.Fatal(err)
		}
		cfg := client.Config{
			Endpoints:globalEndpoints,
			Transport:transport,
			HeaderTimeoutPerRequest:time.Second,
		}
		cli, err := client.New(cfg)
		if err != nil {
			log.Fatal(err)
		}
		globalKapi = client.NewKeysAPI(cli)
	}
	return globalKapi
}

/**
	获取transport
 */
func getTransport() (*http.Transport, error) {
	if globalCaFile == "" {
		globalCaFile = os.Getenv("CA_FILE")
	}
	if globalCertFile == "" {
		globalCertFile = os.Getenv("CERT_FILE")
	}
	if globalKeyFile == "" {
		globalKeyFile = os.Getenv("KEY_FILE")
	}

	tls := transport.TLSInfo{
		CAFile:   globalCaFile,
		CertFile: globalCertFile,
		KeyFile:  globalKeyFile,
	}
	return transport.NewTransport(tls, 30 * time.Second)
}

/**
	某个目录下的所有键
 */
func QueryKeyList(path string) (*client.Response, error) {
	kapi := getKapi()
	opts := &client.GetOptions{}
	opts.Recursive = true
	return kapi.Get(context.Background(), path, opts)
}

/**
	更新某个键的值
 */
func UpdateKey(k, v string) (*client.Response, error) {
	kapi := getKapi()
	return kapi.Update(context.Background(), k, v)
}

/**
	设置某个键的值
 */
func SetKey(k, v string, opts *client.SetOptions) (*client.Response, error) {
	kapi := getKapi()
	return kapi.Set(context.Background(), k, v, opts)
}

/**
	删除某个键
 */
func DeleteKey(k string, opts *client.DeleteOptions) (*client.Response, error) {
	kapi := getKapi()
	return kapi.Delete(context.Background(), k, opts)
}