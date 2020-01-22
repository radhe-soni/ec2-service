package main

import (
	"bufio"
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
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("-> ")
    suffix, _ := reader.ReadString('\n')
	for _, b := range buckets {
		
		if strings.EqualFold(*cfg.Region, bs.GetBucketRegion(*b.Name)) && strings.HasSuffix(*b.Name, suffix) {
			fmt.Println("##################################")
			bucketservice.DisplayBucket(b)
			bs.ListObject(*b.Name)
			bs.DeleteBucket(*b.Name)
		}
        //if strings.end
	}
	
	
	//bucketservice.CreateBucket(s3Client, "hpcl.radhe.dev")
}
