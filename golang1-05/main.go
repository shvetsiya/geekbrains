package main

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v2"

	"github.com/shvetsiya/geekbrains/golang1-05/config"
)

func main() {
	file, err := os.Open("./config.yaml")
	if err != nil {
		log.Fatalf("Can not open config file %v", err)
	}

	defer file.Close()

	conf := new(config.Config)
	err = yaml.NewDecoder(file).Decode(conf)
	if err != nil {
		log.Fatalf("Can not decode yaml file %v", err)
	}

	if !conf.IsValid() {
		log.Fatalf("Config contains not valid fields %v", err)
	}
	fmt.Printf("%v\n", *conf)
}
