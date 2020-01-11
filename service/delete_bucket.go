package bucketservice

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

func deleteEmptyBucket(s3Client *s3.S3, bucket string){
	// Delete the S3 Bucket
	// It must be empty or else the call fails
	deleteBucketInput := s3.DeleteBucketInput{
        Bucket: aws.String(bucket),
	}
    _, err := s3Client.DeleteBucket(&deleteBucketInput)
    if err != nil {
        log.Printf("Unable to delete bucket %q, %v", bucket, err)
    }

    // Wait until bucket is deleted before finishing
    log.Printf("Waiting for bucket %q to be deleted...\n", bucket)

    err = s3Client.WaitUntilBucketNotExists(&s3.HeadBucketInput{
        Bucket: aws.String(bucket),
    })
    if err != nil {
        log.Printf("Error occurred while waiting for bucket to be deleted, %v", bucket)
    }

    log.Printf("Bucket %q successfully deleted\n", bucket)
}