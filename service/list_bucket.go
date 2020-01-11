package bucketservice

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func listBucket(s3Client *s3.S3) {
	result, err := s3Client.ListBuckets(nil)
	if err != nil {
		log.Println("Failed to list buckets", err)
		return
	}

	fmt.Println("Buckets:")
	fmt.Println("Bucket Name \t\t Created On")
	for _, b := range result.Buckets {
		fmt.Printf("* %s \t\t %s\n",
			aws.StringValue(b.Name), aws.TimeValue(b.CreationDate))
	}
}
