package s3

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"log"
	"time"
)

func GenerateSignedURL(bucket string, key string) (string, error) {

	sess, err := session.NewSession(&aws.Config{Region: aws.String("eu-west-2")})

	// Create S3 service client
	svc := s3.New(sess)

	output, err := svc.HeadObject(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	if err != nil {
		fmt.Errorf(err.Error())
	}

	fmt.Print(output.LastModified)
	if output.LastModified == nil {
		return "", errors.New(fmt.Sprintf("object %s/%s not found", bucket, key))
	}

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	urlStr, err := req.Presign(15 * time.Minute)

	if err != nil {
		log.Println("Failed to sign request", err)
	}

	return urlStr, nil
}
