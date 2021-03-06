package config

import (
	"io/ioutil"
	"path/filepath"
	"gopkg.in/yaml.v1"
	"log"
)

type DBconfig struct {
	Alias        string `yaml:"alias"`
	DriverName   string `yaml:"drive_name"`
	DataSource   string `yaml:"data_source"`
	MaxIdleConns int `yaml:"max_idle_conns"`
	MaxOpenConns int `yaml:"max_open_conns"`
	CreateTable  bool `yaml:"create_table"`
}

type ETCDconfig struct {
	EndPoints []string `yaml:"endpoints"`
	BasePath  string   `yaml:"base_path"`
	CaFile    string `yaml:"ca_file"`
	CertFile  string `yaml:"cert_file"`
	KeyFile   string `yaml:"key_file"`
}

/**
	获取数据库配置
 */
func GetDBconfig() *DBconfig {
	filename, _ := filepath.Abs("./config/DBconfig.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
	}

	var c *DBconfig
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}

/**
	获取etcd配置
 */
func GetETCDconfig() *ETCDconfig {
	filename, err := filepath.Abs("./config/ETCDconfig.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}

	var c *ETCDconfig
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		log.Fatal(err)
	}
	return c
}