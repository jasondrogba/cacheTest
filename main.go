package main

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"jasondrogba/alluxio-cacheTest/alluxioTest"
	"jasondrogba/alluxio-cacheTest/ec2test"
	"jasondrogba/alluxio-cacheTest/getArgTest"
	"jasondrogba/alluxio-cacheTest/metricsTest"
	"jasondrogba/alluxio-cacheTest/sshTest"
	"jasondrogba/alluxio-cacheTest/startTest"
	"log"
)

func main() {

	_, _, circle := getArgTest.ParseArgs()
	instanceMap := ec2test.Getec2Instance()

	for count := 3; count < 10; count++ {
		resultRemotesLRU := make([]float64, 0)
		resultRemotesREPLICA := make([]float64, 0)
		resultUFSsLRU := make([]float64, 0)
		resultUFSsREPLICA := make([]float64, 0)
		for i := 0; i < circle; i++ {
			fmt.Println("@@@@@@@@@start Alluxio@@@@@@@@@@:", "LRU")
			startTest.StartTest(instanceMap["Ec2Cluster-default-masters-0"], "LRU")
			fmt.Println("@@@@@@@@@LOAD Alluxio@@@@@@@@@@:", "worker0:1~18,worker1:1~10:worker2:1~5")
			sshTest.SshTest(instanceMap)
			fmt.Println("@@@@@@@@@READ Alluxio@@@@@@@@@@:", count*100)
			alluxioTest.ReadAlluxio(instanceMap["Ec2Cluster-default-masters-0"], count*100)
			fmt.Println("@@@@@@@@@METRIC Alluxio@@@@@@@@@@:")
			resultRemote, resultUFS := metricsTest.BackProcess(instanceMap)
			resultRemotesLRU = append(resultRemotesLRU, resultRemote)
			resultUFSsLRU = append(resultUFSsLRU, resultUFS)
			fmt.Println("@@@@@@@@@@circle@@@@@@@@@@@@@@:", i)
		}
		for i := 0; i < circle; i++ {
			fmt.Println("@@@@@@@@@start Alluxio@@@@@@@@@@:", "REPLICA")
			startTest.StartTest(instanceMap["Ec2Cluster-default-masters-0"], "REPLICA")
			fmt.Println("@@@@@@@@@LOAD Alluxio@@@@@@@@@@:", "worker0:1~18,worker1:1~10:worker2:1~5")
			sshTest.SshTest(instanceMap)
			fmt.Println("@@@@@@@@@READ Alluxio@@@@@@@@@@:", count*100)
			alluxioTest.ReadAlluxio(instanceMap["Ec2Cluster-default-masters-0"], count*100)
			fmt.Println("@@@@@@@@@METRIC Alluxio@@@@@@@@@@:")
			resultRemote, resultUFS := metricsTest.BackProcess(instanceMap)
			resultRemotesREPLICA = append(resultRemotesREPLICA, resultRemote)
			resultUFSsREPLICA = append(resultUFSsREPLICA, resultUFS)
			fmt.Println("@@@@@@@@@@circle@@@@@@@@@@@@@@:", i)
		}
		fmt.Println("******【read count】*********: ", count)
		fmt.Println("resultUFSLRU:", resultUFSsLRU)
		fmt.Println("resultRemoteLRU:", resultRemotesLRU)
		fmt.Println("resultUFSREPLICA:", resultUFSsREPLICA)
	}
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
