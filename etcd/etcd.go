package etcd

import (
	"log"
	"time"

	"github.com/coreos/etcd/client"
	"net/http"
	"os"
	"github.com/coreos/etcd/pkg/transport"
)

var globalEndpoints []string
var globalCaFile string
var globalCertFile string
var globalKeyFile string

var globalKapi client.KeysAPI

func InitConfig(endpoints []string, caFile string, certFile string, keyFile string) {
	globalEndpoints = endpoints
	globalCaFile = caFile
	globalCertFile = certFile
	globalKeyFile = keyFile
}

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
	}
}

func getTransport() (*http.Transport, error) {
	if globalCaFile == "" {
		globalCaFile = os.Getenv("DCMP_CA_FILE")
	}
	if globalCertFile == "" {
		globalCertFile = os.Getenv("DCMP_CERT_FILE")
	}
	if globalKeyFile == "" {
		globalKeyFile = os.Getenv("DCMP_KEY_FILE")
	}

	tls := transport.TLSInfo{
		CAFile:   globalCaFile,
		CertFile: globalCertFile,
		KeyFile:  globalKeyFile,
	}

	return transport.NewTransport(tls, 30*time.Second)
}