package main

import (
	bucketservice "github.com/radhe-soni/s3-bucket-service/service"
)

func main() {
	sess := bucketservice.GetSession()
	bs := bucketservice.NewBucketService(sess)
	bs.ListBucket()
	//bucketservice.CreateBucket(s3Client, "hpcl.radhe.dev")
}
