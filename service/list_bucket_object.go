package bucketservice

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func listObject(s3Client *s3.S3, bucket string) {
	contents := getAllObjects(s3Client,bucket)
    for _, item := range contents {
        fmt.Println("Name:         ", *item.Key)
        fmt.Println("Last modified:", *item.LastModified)
        fmt.Println("Size:         ", *item.Size)
        fmt.Println("Storage class:", *item.StorageClass)
        fmt.Println("--------------------------------------------")
    }

    fmt.Println("Found", len(contents), "items in bucket", bucket)
}

func getAllObjects(s3Client *s3.S3, bucket string) []*s3.Object{
// Get the list of items
    resp, err := s3Client.ListObjectsV2(&s3.ListObjectsV2Input{Bucket: aws.String(bucket)})
    if err != nil {
        log.Printf("Unable to list items in bucket %q, %v", bucket, err)
	}
	return resp.Contents
}