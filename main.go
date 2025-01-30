package main

import (
	"go-web/common/job"
	"go-web/common/log"
	"go-web/common/router"
	"net/http"
)

func main() {
	logger := log.GetLogger()
	mux := router.Routers()
	job.StartJob()
	logger.Info("start server on port: 8080")
	http.ListenAndServe(":8080", mux)
}
