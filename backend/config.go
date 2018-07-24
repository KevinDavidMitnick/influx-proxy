// Copyright 2016 Eleme. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package backend

import (
	"encoding/json"
	"log"
	"os"
	"sync"
)

const (
	VERSION = "1.1"
)

type NodeConfig struct {
	ListenAddr   string
	DB           string
	Zone         string
	Nexts        string
	Interval     int
	IdleTimeout  int
	WriteTracing int
	QueryTracing int
}

type BackendConfig struct {
	URL             string
	DB              string
	Zone            string
	Interval        int
	Timeout         int
	TimeoutQuery    int
	MaxRowLimit     int
	CheckInterval   int
	RewriteInterval int
	WriteOnly       int
}

type ConfigSource struct {
	Backends map[string]*BackendConfig `json:"backends"`
	Node     *NodeConfig               `json:"node"`
}

var (
	config *ConfigSource
	lock   = new(sync.RWMutex)
)

func Config() *ConfigSource {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	file, err := os.Open(cfg)
	if err != nil {
		log.Fatalln("config file:", cfg, "open failed.")
	}
	defer file.Close()

	var c ConfigSource
	err = json.NewDecoder(file).Decode(&c)
	if err != nil {
		log.Fatalln("config file:", cfg, "decode  failed.")
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c

	log.Println("read config file:", cfg, "successfully")
}
