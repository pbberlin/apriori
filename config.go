package main

import (
	"encoding/json"
	"io"
	"log"
	"os"
	"path"
)

type Config struct {
	AppName string `json:"app_name"`
	Carts   TCarts `json:"carts_with_meta"`
}

var Conf Config

func LoadConfig() io.Reader {
	fName := "config.json"
	fPath := path.Join(".", fName)
	file, err := os.Open(fPath)
	if err != nil {
		log.Fatalf("could not open: %v; ", err)
	}
	return file
}

func init() {
	log.SetFlags(log.Lshortfile)
	fileReader := LoadConfig()
	decoder := json.NewDecoder(fileReader)
	err := decoder.Decode(&Conf)
	if err != nil {
		log.Fatalf("could not decode json config: %v; ", err)
	}

	return
	formatted, err := json.MarshalIndent(Conf, " ", "\t")
	if err != nil {
		log.Fatalf("could not format json config: %v; ", err)
	}
	log.Printf("\n%#s", string(formatted))
}
