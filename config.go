package main

import (
	"io/ioutil"
	"path/filepath"
	"gopkg.in/yaml.v1"
)

type DBconfig struct {

}

type ETCDconfig struct {

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