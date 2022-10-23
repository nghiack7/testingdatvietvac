package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"testing.com/datviecvac/pkg/cloud"
	"testing.com/datviecvac/src/api"
)

type Body struct {
	Data string `json:"data"`
	Name string `json:"name"`
}

func TestUploadDataApi(t *testing.T) {
	gin.SetMode(gin.TestMode)

	mockResponse := `{"message":"Accepted uploaded the file"}`
	//create s3 client
	config := cloud.Config{
		Address: "http://localhost:4566",
		Region:  "eu-west-1",
		Profile: "localstack",
		ID:      "test",
		Secret:  "test"}
	newssess := cloud.ConnectAWS(config)
	newS3 := cloud.NewS3(newssess, 10*time.Second)
	r := gin.Default()
	srv := &api.Server{
		Client: newS3,
		Router: r,
	}
	r.POST("/user/batch", srv.UploadData)
	mockBody := Body{
		Data: "data using in test when create body",
		Name: "this is body using for test",
	}
	jsonVal, _ := json.Marshal(mockBody)
	req, err := http.NewRequest(http.MethodPost, "/user/batch", bytes.NewBuffer(jsonVal))
	if err != nil {
		//TODO
	}
	w := httptest.NewRecorder()

	r.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	responseData, _ := ioutil.ReadAll(w.Body)
	//assert value response
	assert.Equal(t, mockResponse, string(responseData))
}
