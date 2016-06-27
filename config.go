package main

import (
	"encoding/json"
	"io/ioutil"
)

type Configure struct {
	ListenAddr  string
	SecertKey   string
	MaxMemory   int
	JarPath     string
	WorkDir     string
	StopCommand string
	AllowRemoteExec  bool
}

var cfg Configure

func loadConfigure(jsonPath string) {
	jsonStr, err := ioutil.ReadFile(jsonPath)
	checkError(err)
	err = json.Unmarshal(jsonStr, &cfg)
	checkError(err)
}
