package bucketservice

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/service/s3"
)

func listBucket(s3Client *s3.S3) {
	buckets := getAllBuckets(s3Client)

	fmt.Println("Buckets:")
	fmt.Println("Bucket Name \t\t Created On")
	for _, b := range buckets {
		DisplayBucket(b)
	}
}

func getAllBuckets(s3Client *s3.S3) []*s3.Bucket {
	var buckets []*s3.Bucket
	result, err := s3Client.ListBuckets(nil)
	if err != nil {
		log.Println("Failed to list buckets", err)
		return buckets
	}
	buckets = result.Buckets
	return buckets
}
