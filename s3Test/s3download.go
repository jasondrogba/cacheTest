package s3Test

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"os"
	"time"
)

func DownloadS3Bucket() {
	bucket := "alluxio-tpch100"
	item := "1.txt"

	file, err := os.Create(item)
	if err != nil {
		fmt.Println(err)
	}

	startTime := time.Now()

	sess, _ := session.NewSession(&aws.Config{Region: aws.String("us-east-1")})
	downloader := s3manager.NewDownloader(sess)

	numBytes, err := downloader.Download(file,
		&s3.GetObjectInput{
			Bucket: aws.String(bucket),
			Key:    aws.String("cacheTest/" + item),
		})
	if err != nil {
		fmt.Println(err)
	}
	elapsed := time.Since(startTime).Seconds()
	downloadSpeed := float64(numBytes) / elapsed / 1024 / 1024 // 带宽以KB/s为单位

	fmt.Println("Downloaded", file.Name(), numBytes, "bytes", " 下载带宽：", downloadSpeed, "MB/s")
}
