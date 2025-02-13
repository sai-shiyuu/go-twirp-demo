package main

import (
	"flag"
	"fmt"
	"go-web/common/config"
	"go-web/common/job"
	"go-web/common/log"
	"go-web/common/router"
	"net/http"
)

func main() {
	env := flag.String("env", "dev", "environment: dev or prod")
	flag.Parse()
	logger := log.GetLogger()
	c, err := config.LoadConfig("config.toml")
	if err != nil {
		logger.Fatalf("Error loading config: %s", err)
	}
	if *env == "prod" {
		config.Configs = &c.Prod
	} else {
		config.Configs = &c.Dev
	}
	logger.Println("Loaded config:", *env, config.Configs)
	mux := router.Routers()
	job.StartJob()
	logger.Infof("start server on port: %d", config.Configs.Server.Port)
	http.ListenAndServe(fmt.Sprintf(":%d", config.Configs.Server.Port), mux)
}
