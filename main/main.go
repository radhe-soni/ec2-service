package main

import (
	"fmt"
	"strings"
	bucketservice "github.com/radhe-soni/s3-bucket-service/service"
)

func main() {
    fmt.Println("SDS : Simply delete Service")
	sess := bucketservice.GetSession()
	 cfg := sess.Config
	bs := bucketservice.NewBucketService(sess)
	buckets := bs.GetAllBuckets()
	fmt.Println("Enter the bucket name suffix to delete -> ")
	var suffix string
    fmt.Scanln(&suffix)
	for _, b := range buckets {
        var name = *b.Name
		if strings.EqualFold(*cfg.Region, bs.GetBucketRegion(name)) && strings.HasSuffix(name, suffix) {
			fmt.Println("##################################")
			bucketservice.DisplayBucket(b)
			bs.ListObject(name)
			bs.DeleteBucket(name)
		}
	}
	
	
	//bucketservice.CreateBucket(s3Client, "hpcl.radhe.dev")
}