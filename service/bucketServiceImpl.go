package bucketservice

import (
	"fmt"
	"log"
	"os"
	"strings"
	"context"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)
//const s3Suffix = ".radhe.dev"
const s3Suffix = ""
type bucketService struct {
	awsSession *session.Session
	s3Client   *s3.S3
	s3Suffix   string
}

//NewBucketService : Returns new BucketService
func NewBucketService(awsSession *session.Session) BucketService {
	s3Client := s3.New(awsSession)
	bs := bucketService{awsSession: awsSession, s3Client: s3Client, s3Suffix: s3Suffix}
	return bs
}

//GetSession : returns an instance of aws session
func GetSession() *session.Session {
	config := aws.Config{Region: aws.String("ap-south-1")}
	sess, err := session.NewSession(&config)
	if err != nil {
		log.Fatal("Error occured while getting session ", err)
	}
	log.Println("aws session obtained")
	return sess
}

/*ListBucket : Takes an instance of aws/aws-sdk-go/service/s3.S3
and displays list of all buckets*/
func (bs bucketService) ListBucket() {
	listBucket(bs.s3Client)
}
func (bs bucketService) GetAllBuckets() []*s3.Bucket {
	return getAllBuckets(bs.s3Client)
}
func (bs bucketService) CreateBucket(bucket string) string {
	bucket = strings.ToLower(bucket + bs.s3Suffix)
	createBucket(bs.s3Client, bucket)
	return bucket
}

func (bs bucketService) ListObject(bucket string) {
	listObject(bs.s3Client, bucket)
}
func (bs bucketService) DeleteBucket(bucket string) string {
	bucket = strings.ToLower(bucket + bs.s3Suffix)
	deleteAllObject(bs.s3Client, bucket)
	deleteEmptyBucket(bs.s3Client, bucket)
	return bucket
}

// GetBucketRegion : gives bucket region name
func (bs bucketService) GetBucketRegion(bucket string) string {
	ctx := context.Background()
	region, err := s3manager.GetBucketRegion(ctx, bs.awsSession, bucket, "")
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == "NotFound" {
			fmt.Fprintf(os.Stderr, "unable to find bucket %s's region not found\n", bucket)
		}
		return ""
	}
	fmt.Printf("Bucket %s is in %s region\n", bucket, region)
	return region
}