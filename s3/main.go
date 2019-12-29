package s3

import (
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

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

func createBucket() {
	fmt.Printf("creating bucket %s\n...", myBucket)
	result, err := s3Instance.CreateBucket(&s3.CreateBucketInput{
		Bucket: aws.String(myBucket),
	})
	HandlerErr(err)
	fmt.Println(result)
}

// Upload ... accepts an array of bytes.
func Upload(data string) {
	fmt.Println("upload method was called")
	bucketExists := CheckBucketExists(myBucket, s3Instance)
	if !bucketExists {
		createBucket()
	}

	// myBody := strings.NewReader("this is another test")
	uploader := s3manager.NewUploader(sess)

	_, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myKey),
		Body:   strings.NewReader(data),
	})
	HandlerErr(err)

	fmt.Println("object added")
}

type writer struct {
	data []byte
}

func (w writer) WriteAt(p []byte, off int64) (n int, err error) {
	fmt.Println(string(p))
	return len(p), nil
}

func Download() {
	downloader := s3manager.NewDownloader(sess)
	w := writer{}
	result, err := downloader.Download(w, &s3.GetObjectInput{
		Bucket: aws.String(myBucket),
		Key:    aws.String(myKey),
	})
	HandlerErr(err)
	fmt.Println(result)
}
