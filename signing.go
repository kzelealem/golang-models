package models

import (
	"net/url"
	"os"
	"strings"
	"sync"
	"time"

	"github.com/IBM/ibm-cos-sdk-go/aws"
	"github.com/IBM/ibm-cos-sdk-go/aws/credentials"
	"github.com/IBM/ibm-cos-sdk-go/aws/session"
	"github.com/IBM/ibm-cos-sdk-go/service/s3"
)

var (
	once    sync.Once
	client  *s3.S3
	errInit error
)

func getClient() (*s3.S3, error) {
	once.Do(func() {
		region := os.Getenv("STORAGE_REGION")
		accessKeyID := os.Getenv("STORAGE_READ_ACCESS_KEY_ID")
		secretAccessKey := os.Getenv("STORAGE_READ_SECRET_ACCESS_KEY")
		endpoint := os.Getenv("STORAGE_ENDPOINT")

		sess := session.Must(session.NewSession(&aws.Config{
			Credentials: credentials.NewStaticCredentials(accessKeyID, secretAccessKey, ""),
			Endpoint:    aws.String(endpoint),
			Region:      aws.String(region), // Choose the appropriate region
		}))
		credentials.NewStaticCredentialsFromCreds(credentials.Value{})

		client = s3.New(sess)
	})
	return client, errInit
}

func GenerateSignedURL(fullURL string) (string, error) {
	client, err := getClient()
	if err != nil {
		return "", err
	}

	// Parse the full URL
	parsedURL, err := url.Parse(fullURL)
	if err != nil {
		return "", err
	}

	// Extract bucket name and object key
	// Assuming the URL format is: https://[BUCKET_NAME].s3.[REGION].amazonaws.com/[OBJECT_KEY]
	parts := strings.SplitN(parsedURL.Host, ".", 2)
	if len(parts) < 2 {
		return "", nil
	}
	bucketName := parts[0]
	objectKey := strings.TrimLeft(parsedURL.Path, "/")

	req, _ := client.GetObjectRequest(&s3.GetObjectInput{
		Bucket: aws.String(bucketName),
		Key:    aws.String(objectKey),
	})

	return req.Presign(15 * time.Minute) // Adjust the expiration time as needed
}
