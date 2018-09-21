package main

import (
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
)

type Yaml struct {
	Name   string
	Domain string
	Links  []string
	Tags   Tags
}

type Tags struct {
	Title    string
	Code     string
	Store    string
	Internet string
	Sales    string
	Price    string
}

/**
Method to load link file and then return it into array
*/
func LoadFile(filename string) (result Yaml) {

	var conf Yaml

	if fileValidate(filename) {
		yamlFile, _ := os.Open(filename)
		buf, _ := ioutil.ReadAll(yamlFile)
		yaml.Unmarshal(buf, &conf)

		if len(conf.Links) < 1 {
			log.Fatal("The file has no links")
			os.Exit(1)
		}
	}

	return conf
}

/**
Validate the filename exist returning a boolean value
*/
func fileValidate(filename string) bool {
	if _, err := os.Stat(filename); err == nil {
		return true
	}

	return false
}
