package bucketservice

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

//BucketService : Used for S3 bucket crud operations
type BucketService interface {
	ListBucket()
	CreateBucket(bucket string) string
}

type bucketService struct {
	awsSession *session.Session
	s3Client   *s3.S3
	s3Suffix   string
}

//NewBucketService : Returns new BucketService
func NewBucketService(awsSession *session.Session) *bucketService {
	s3Client := s3.New(awsSession)
	bs := bucketService{awsSession, s3Client, "carrier.radhe"}
	return &bs
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

func (bs bucketService) CreateBucket(bucket string) string {
	bucket = strings.ToLower(bucket + bs.s3Suffix)
	createBucket(bs.s3Client, bucket)
	return bucket
}
