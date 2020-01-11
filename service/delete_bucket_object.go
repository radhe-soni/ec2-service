package bucketservice
import (
	"fmt"
	"log"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func deleteAllObject(s3Client *s3.S3, bucket string){
	iter := s3manager.NewDeleteListIterator(s3Client, &s3.ListObjectsInput{
        Bucket: aws.String(bucket),
    })

    // Traverse iterator deleting each object
    if err := s3manager.NewBatchDeleteWithClient(s3Client).Delete(aws.BackgroundContext(), iter); err != nil {
        log.Printf("Unable to delete objects from bucket %q, %v", bucket, err)
    }

    fmt.Printf("Deleted object(s) from bucket: %s", bucket)
}