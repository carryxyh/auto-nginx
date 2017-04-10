package config

import (
	"io/ioutil"
	"path/filepath"
	"gopkg.in/yaml.v1"
	"log"
)

type DBconfig struct {
	Alias      string `yaml:"alias"`
	DriverName string `yaml:"drive_name"`
	DataSource string `yaml:"data_source"`
	Conns      []int `yaml:"conns"`
}

type ETCDconfig struct {
	Listen    string `yaml:"listen"`
	EndPoints []string `yaml:"endpoints"`
	BasePath  string   `yaml:"base_path"`
}

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