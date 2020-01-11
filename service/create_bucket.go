package bucketservice

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func createBucket(s3Client *s3.S3, bucket string) {
	request := s3.CreateBucketInput{
		Bucket: aws.String(bucket),
	}
	_, err := s3Client.CreateBucket(&request)
	if err != nil {
		log.Printf("Unable to create bucket %q, %v", bucket, err)
	}

	// Wait until bucket is created before finishing
	log.Printf("Waiting for bucket %q to be created...\n", bucket)

	err = s3Client.WaitUntilBucketExists(&s3.HeadBucketInput{
		Bucket: aws.String(bucket),
	})
	if err != nil {
		log.Printf("Error occurred while waiting for bucket to be created, %v", bucket)
	}

	fmt.Printf("Bucket %q successfully created\n", bucket)
}
