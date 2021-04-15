package server 

import (
	"os"
	"fmt"
	"time"
	"context"
	"io"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	bucket = "BUCKET"
	key    = "KEY"
)

var (
	timeout = time.Duration(time.Minute * 1)
)

func getFile(path string) (io.ReadSeeker, error) {
	data, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %s", err.Error())
	}
	return data, nil
}

func upload(filePath string) (string, error) {
	fileObject, err := getFile(filePath)
	if err != nil {
		return "", err
	}

	sess := session.Must(session.NewSession())
	service := s3.New(sess)

	ctx := context.Background()
  	var cancelFn func()
  	if timeout > 0 {
  		ctx, cancelFn = context.WithTimeout(ctx, timeout)
  	}

  	if cancelFn != nil {
  		defer cancelFn()
	}

	_, err = service.PutObjectWithContext(ctx, &s3.PutObjectInput{
  		Bucket: aws.String(bucket),
  		Key:    aws.String(key),
  		Body:   fileObject,
  	})

  	if err != nil {
  		return "", err
  	}

  	return fmt.Sprintf("%s/%s\n", bucket, key), nil
}