package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"jasondrogba/cacheTest/alluxioTest"
	"jasondrogba/cacheTest/ec2test"
	"jasondrogba/cacheTest/metricsTest"
	"jasondrogba/cacheTest/sshTest"
	"jasondrogba/cacheTest/startTest"
	"log"
	"os"
)

func main() {
	instanceMap := ec2test.Getec2Instance()
	sshTest.SshTest(instanceMap)
	alluxioTest.ReadAlluxio(instanceMap["Ec2Cluster-default-masters-0"])
	startTest.Starttest(instanceMap["Ec2Cluster-default-masters-0"])
	metricsTest.BackProcess()

}

// 写一个脚本程序，可以远程ssh连接到aws服务器，在aws中执行command，然后将结果返回到本地
func s3list() {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}
	// Create an Amazon S3 service client

	client := s3.NewFromConfig(cfg)

	// Get the list of buckets
	// Get the first page of results for ListObjectsV2 for a bucket
	output, err := client.ListObjectsV2(context.TODO(), &s3.ListObjectsV2Input{
		Bucket: aws.String("alluxio-tpch100"),
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Println("first page results:")
	for _, object := range output.Contents {
		log.Printf("key=%s size=%d", aws.ToString(object.Key), object.Size)
	}

}

func ec2connect2() {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		os.Exit(1)
	}

	// Create EC2 client
	client := ec2instanceconnect.NewFromConfig(cfg)

	output, err := client.SendSSHPublicKey(context.TODO(), &ec2instanceconnect.SendSSHPublicKeyInput{
		AvailabilityZone: aws.String("us-east-1a"),
		InstanceId:       aws.String("i-078d9ffc684629fa3"),
		InstanceOSUser:   aws.String("ec2-user"),
		SSHPublicKey:     aws.String("ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABgQDEwu/CjCX+H1xzfFG9DRthQVNNrio8dB+6o/paSwuBVZJM9plv3lhDc16beHRqzd5aoraeTUlkEaxTS9IDXD61ht5Fx3j+NLQY0u5MmCwOABqYrmLTUXT3Y7Lj5YZNQOhKnignsbSzdKQekCPc20j2FmkCZjS+FgCV/bi/eYT8BGYkBbogI4dPSU45n99b62F12EUErHyn8jxkwtjwTJEW+k12dDVelFIjHbswryDjP5RmwxDxStSxVeZvkL3r2cg1ifYPYnaNXpUcxK3wD8RjKZu3X41dEr66aRvo2GCDIUEWn7aquXC0mArnUJfAVZubwWLPK8V8JNlSGsWa2swj1T8N+/SyecWRD9AWEuKUXHPIvRsUW9Wbp4Puryx0edqUK9gJ+VlPcprBY9hZColTXCDTtXKan9Ud5oZqJxORCF2DRWccJfqCuZePNVdln3mazy8I02+//nVGSKqx1w6hipoP5ewEYmWn65lwiJY+AGLyLqwrh3hgjbl2s2k/I5E= sunbury@sunburydeMBP"), //这里填写你的公钥
	})
	if err != nil {
		fmt.Println("Error creating session:", err)
		os.Exit(1)
	}
	fmt.Print(output)

}
