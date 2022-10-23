package handlers

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"sync"

	"testing.com/datviecvac/pkg/cloud"
)

const (
	AWS_S3_BUCKET = "testdv"
)

func HandlerUpload(data []byte) {

	fileName := fmt.Sprintf("upload-*.%s", "json")
	// Create a temporary file with a dir folder
	tempFile, err := ioutil.TempFile("upload", fileName)

	if err != nil {
		fmt.Println(err)
	}

	defer tempFile.Close()
	tempFile.Write(data)
}

func HandlerUploadToS3(data []byte, client cloud.BucketClient, wg *sync.WaitGroup) {
	defer (*wg).Done()
	ctx := context.Background()
	err := client.Create(ctx, AWS_S3_BUCKET)
	if err != nil {
		fmt.Println("Can't create bucket with err: ", err)
	}
	fileName := fmt.Sprintf("upload-*.%s", "json")
	reader := bytes.NewReader(data)
	str, err := client.UploadObject(ctx, AWS_S3_BUCKET, fileName, reader)
	if err != nil {
		fmt.Println("err update, err= ", err)
	}
	fmt.Println(str)
}
