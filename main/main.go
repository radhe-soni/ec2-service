package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	bucketservice "github.com/radhe-soni/s3-bucket-service/service"
)

func main() {
	if len(os.Args) != 2 {
		log.Printf("bucket name required\nUsage: %s bucket_name", os.Args[0])
	}
	sess := bucketservice.GetSession()
	cfg := sess.Config
	bs := bucketservice.NewBucketService(sess)
	buckets := bs.GetAllBuckets()
	
	for _, b := range buckets {
		fmt.Println("##################################")
		bucketservice.DisplayBucket(b)
		if strings.EqualFold(*cfg.Region, bs.GetBucketRegion(*b.Name)) {
			bs.ListObject(*b.Name)
		}
	}
	bs.DeleteBucket(os.Args[1])
	//bucketservice.CreateBucket(s3Client, "hpcl.radhe.dev")
}
