package ec2test

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2"
	"os"
)

func Ec2command() {
	// Load AWS configuration
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		fmt.Println("Error loading AWS config:", err)
		os.Exit(1)
	}

	instanceID := "i-078d9ffc684629fa3"

	ec2client := ec2.NewFromConfig(cfg)
	describeInstancesInput := &ec2.DescribeInstancesInput{
		InstanceIds: []string{instanceID},
	}

	result, err := ec2client.DescribeInstances(context.Background(), describeInstancesInput)
	if err != nil {
		fmt.Println("Error describing instances:", err)
		os.Exit(1)
	}
	fmt.Println(result.Reservations[0].Instances[0].PublicIpAddress)

}
