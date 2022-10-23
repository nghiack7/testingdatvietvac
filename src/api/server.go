package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"testing.com/datviecvac/pkg/cloud"
	"testing.com/datviecvac/src/handlers"
	"testing.com/datviecvac/src/middleware"
)

type Server struct {
	Client cloud.BucketClient
	Router *gin.Engine
}

// Create New server using S3
func NewServer(client cloud.BucketClient) *Server {
	server := &Server{Client: client}
	server.initRoute()
	return server
}
func (server *Server) initRoute() {
	router := gin.Default()
	//handle middle ware add limit 10kb for body
	router.Use(middleware.BodySizeMaxMiddleWare())
	//init upload router api
	router.POST("/user/batch", server.UploadData)
	server.Router = router
}
func (server *Server) Start(address string) error {
	return server.Router.Run(address)
}

func (server *Server) UploadData(c *gin.Context) {
	var wg sync.WaitGroup
	data, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		fmt.Errorf("Can't get file from Request, err= %v", err)
	}
	//upload to folder upload source
	go handlers.HandlerUpload(data)

	//Upload to s3 store
	wg.Add(1)
	go handlers.HandlerUploadToS3(data, server.Client, &wg)
	wg.Wait()
	c.JSON(http.StatusCreated, gin.H{
		"message": "Accepted uploaded the file",
	})
}
