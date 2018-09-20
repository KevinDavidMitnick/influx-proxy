// Copyright 2016 Eleme. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package main

import (
	"errors"
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/influx-proxy/backend"
	lumberjack "gopkg.in/natefinch/lumberjack.v2"
)

var (
	ErrConfig   = errors.New("config parse error")
	ConfigFile  string
	LogFilePath string
)

func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)

	flag.StringVar(&LogFilePath, "o", "/var/log/influx-proxy.log", "output file")
	flag.StringVar(&ConfigFile, "c", "cfg.json", "config file")
	flag.Parse()
}

func initLog() {
	if LogFilePath == "" {
		log.SetOutput(os.Stdout)
	} else {
		log.SetOutput(&lumberjack.Logger{
			Filename:   LogFilePath,
			MaxSize:    100,
			MaxBackups: 5,
			MaxAge:     7,
		})
	}
}

func main() {
	initLog()

	var err error
	backend.ParseConfig(ConfigFile)

	cs := backend.Config()
	ic := backend.NewInfluxCluster(cs, cs.Node)
	ic.LoadConfig()

	mux := http.NewServeMux()
	NewHttpService(ic, cs.Node.DB).Register(mux)

	log.Printf("http service start.")
	server := &http.Server{
		Addr:        cs.Node.ListenAddr,
		Handler:     mux,
		IdleTimeout: time.Duration(cs.Node.IdleTimeout) * time.Second,
	}
	if cs.Node.IdleTimeout <= 0 {
		server.IdleTimeout = 10 * time.Second
	}
	err = server.ListenAndServe()
	if err != nil {
		log.Print(err)
		return
	}
}
