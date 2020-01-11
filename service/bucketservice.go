package bucketservice

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

//BucketService : Used for S3 bucket crud operations
type BucketService interface {
	ListBucket()
	GetAllBuckets() []*s3.Bucket
	ListObject(bucket string)
	CreateBucket(bucket string) string
	DeleteBucket(bucket string) string
	GetBucketRegion(bucket string) string
}

// DisplayBucket : utility method to display basic bucket information
func DisplayBucket(bucket *s3.Bucket) {
	fmt.Printf("* %s \t\t %s\n",
		aws.StringValue(bucket.Name), aws.TimeValue(bucket.CreationDate))
}


