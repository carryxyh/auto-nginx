package main

import (
	"io/ioutil"
	"path/filepath"
	"gopkg.in/yaml.v1"
)

type DBconfig struct {
	Listen    string `yaml:"listen"`
	EndPoints []string `yaml:"endpoints"`
	BasePath  string   `yaml:"base_path"`
}

type ETCDconfig struct {
	Alias      string `yaml:"alias"`
	DriverName string `yaml:"drive_name"`
	DataSource string `yaml:"data_source"`
	Conns      []int `yaml:"conns"`
}

func getDBconfig() *DBconfig {
	filename, _ := filepath.Abs("./config/DBconfig.yml")
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var c *DBconfig
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		panic(err)
	}
	return c
}

func getETCDconfig() *ETCDconfig {
	filename, err := filepath.Abs("./config/ETCDconfig.yml")
	yamlFile, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	var c *ETCDconfig
	err = yaml.Unmarshal(yamlFile, &c)
	if err != nil {
		panic(err)
	}
	return c
}