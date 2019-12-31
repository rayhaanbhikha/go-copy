package s3

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// HandlerErr ... logs error
func HandlerErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

// CheckBucketExists ... checks if a bucket exists
func CheckBucketExists(bucketName string, s3Instance *s3.S3) bool {
	result, err := s3Instance.ListBuckets(&s3.ListBucketsInput{})
	HandlerErr(err)
	for _, bucket := range result.Buckets {
		if aws.StringValue(bucket.Name) == bucketName {
			return true
		}
	}
	return false
}
