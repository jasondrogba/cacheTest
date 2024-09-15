package ec2test

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ec2instanceconnect"
	"os"
)

func Ec2connect() {
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
