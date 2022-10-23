package main

import (
	"errors"
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"testing.com/datviecvac/pkg/cloud"
	"testing.com/datviecvac/pkg/utils"
	"testing.com/datviecvac/src/api"
	"testing.com/datviecvac/src/cors"
)

var router *gin.Engine
var newS3 cloud.BucketClient

func init() {
	config := cloud.Config{
		Address: "http://localstack:4566",
		Region:  "eu-west-1",
		Profile: "localstack",
		ID:      "test",
		Secret:  "test"}
	newssess := cloud.ConnectAWS(config)
	newS3 = cloud.NewS3(newssess, 50000*time.Second)
}

func main() {
	server := api.NewServer(newS3)
	err := server.Start("0.0.0.0:8080")
	if err != nil {
		log.Fatal("cannot start server:", err)
	}
}

func noMethodHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.Method == "OPTIONS" {
			cors.CorsMiddleware()(c)
			return
		}
		err := errors.New("The request method is not allowed: " + c.Request.URL.String())
		utils.BuildResponseWithGinError(err, 404, c)
	}
}

func noRouteHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		// For cors
		if c.Request.Method == "OPTIONS" {
			cors.CorsMiddleware()(c)
			return
		}

		err := errors.New("Not found handler for the request path: " + c.Request.URL.String())
		utils.BuildResponseWithGinError(err, 404, c)
	}
}
