package services

import (
	"drink-api/config"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetPresignURL get a presign URL by given key.
// use PUT method to upload file to s3 bucket.
func GetPresignURL(key string) (url string, err error) {
	cred := credentials.NewStaticCredentials(
		config.MustGet("AWS_ACCESS_KEY_ID"),
		config.MustGet("AWS_SECRET_ACCESS_KEY"),
		"",
	)
	sess := session.Must(session.NewSession())

	s3Svc := s3.New(sess, &aws.Config{
		Region:      aws.String("ap-northeast-1"),
		Credentials: cred,
	})

	req, _ := s3Svc.PutObjectRequest(&s3.PutObjectInput{
		Bucket: aws.String(config.MustGet("AWS_BUCKET")),
		Key:    aws.String(key),
		ACL:    aws.String("public-read"),
	})

	url, err = req.Presign(time.Minute * 15)

	if err != nil {
		return "", err
	}

	return url, nil
}
