package s3

import (
	"fmt"
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func handlerErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

const (
	region   = "eu-west-2"
	myBucket = "go-copy"
	myKey    = "first-example-copy"
)

var (
	config = &aws.Config{
		Region: aws.String(region),
	}
	sess       = session.Must(session.NewSession(config))
	s3Instance = s3.New(sess)
)

func checkBucketExists() bool {
	result, err := s3Instance.ListBuckets(&s3.ListBucketsInput{})
	handlerErr(err)
	for _, bucket := range result.Buckets {
		if aws.StringValue(bucket.Name) == myBucket {
			return true
		}
	}
	return false
}

func createBucket() {
	fmt.Printf("creating bucket %s\n...", myBucket)
	result, err := s3Instance.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(myBucket),
	})
	handlerErr(err)
	fmt.Println(result)
}

// Upload ... accepts an array of bytes.
func Upload(data []byte) {
	fmt.Println("upload method was called")
	bucketExists := checkBucketExists()
	if !bucketExists {
		createBucket()
	}

	myBody := strings.NewReader("this is another test")
	uploader := s3manager.NewUploader(sess)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myKey),
		Body:   myBody,
	})
	handlerErr(err)

	fmt.Println("object added")
}
