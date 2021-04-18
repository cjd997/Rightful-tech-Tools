package server

import (
	"fmt"
	"io"
	"os"
	"time"
	"context"

	"github.com/cjd997/Rightful-tech-Tools/chart"
	"github.com/gofiber/fiber/v2"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

func (s *Server) handler(ctx *fiber.Ctx) error {
	var req []chart.Request
	err := ctx.BodyParser(&req)
	if err != nil {
		return fmt.Errorf("error parsing request data: %s", err.Error())
	}

	filePath, err := chart.Generate(req, s.config.Server.FilesFolder)
	if err != nil {
		return err
	}

	// upload file to remote storage
	remoteFilePath, err := s.upload(filePath)
	if err != nil {
		return err
	}

	// delete file locally 
	os.Remove(remoteFilePath)

	// TODO send remoteFilePath as response

	return nil
}

func (s *Server) upload(filePath string) (string, error) {
	fileData, err := getFile(filePath)
	if err != nil {
		return "", err
	}

	sess := session.Must(session.NewSession())
	service := s3.New(sess)

	ctx := context.Background()
	var cancelFn func()

	timeout := time.Duration(time.Second * time.Duration(s.config.Upload.Timeout))
	if timeout > 0 {
		ctx, cancelFn = context.WithTimeout(ctx, timeout)
	}

	if cancelFn != nil {
		defer cancelFn()
	}

	_, err = service.PutObjectWithContext(ctx, &s3.PutObjectInput{
		Bucket: aws.String(s.config.Upload.Bucket),
		Key:    aws.String(s.config.Upload.Key),
		Body:   fileData,
	})

	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s/%s\n", s.config.Upload.Bucket, s.config.Upload.Key), nil
}

func getFile(path string) (io.ReadSeeker, error) {
	data, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error reading file: %s", err.Error())
	}
	return data, nil
}
