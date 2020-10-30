package s3

import (
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"time"
)

func GenerateSignedURL(bucket string, key string) (string, error) {

	sess, err := session.NewSession(&aws.Config{Region: aws.String("eu-west-2")})
	if err != nil {
		return "", err
	}
	svc := s3.New(sess)

	err = verifyObjectExists(svc, bucket, key)
	if err != nil {
		return "", err
	}

	req, _ := svc.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	urlStr, err := req.Presign(15 * time.Minute)
	if err != nil {
		return "", err
	}

	return urlStr, nil
}

func verifyObjectExists(svc *s3.S3, bucket string, key string) error {
	request, output := svc.HeadObjectRequest(&s3.HeadObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})

	err := request.Send()
	if err != nil {
		return err
	}
	if output.LastModified == nil {
		return errors.New(fmt.Sprintf("object %s/%s not found", bucket, key))
	}

	return nil
}
