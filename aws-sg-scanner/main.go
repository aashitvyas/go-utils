package main

import (
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func main() {
	// Initialize a session in us-west-2 that the SDK will use to load
	// credentials from the shared credentials file ~/.aws/credentials.
	sess, err := session.NewSession(&aws.Config{
		Region: aws.String("us-west-2")},
	)

	// Create an EC2 service client.
	svc := ec2.New(sess)

	regions, err := svc.DescribeRegions(&ec2.DescribeRegionsInput{})
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok {
			switch aerr.Code() {
			default:
				fmt.Println(aerr.Error())
			}
		} else {
			// Print the error, cast err to awserr.Error to get the Code and
			// Message from an error.
			fmt.Println(err.Error())
		}
		return
	}

	for _, region := range regions.Regions {
		fmt.Println("Scanning Security Groups in:- ", *region.RegionName)
		// Retrieve the security group descriptions
		sess, err := session.NewSession(&aws.Config{
			Region: aws.String(*region.RegionName)},
		)

		// Create an EC2 service client.
		svc := ec2.New(sess)

		result, err := svc.DescribeSecurityGroups(&ec2.DescribeSecurityGroupsInput{})
		if err != nil {
			if aerr, ok := err.(awserr.Error); ok {
				switch aerr.Code() {
				case "InvalidGroupId.Malformed":
					fallthrough
				case "InvalidGroup.NotFound":
					exitErrorf("%s.", aerr.Message())
				}
			}
			exitErrorf("Unable to get descriptions for security groups, %v", err)
		}
		for _, group := range result.SecurityGroups {
			for _, ippermission := range group.IpPermissions {

				if ippermission.FromPort != nil && *ippermission.FromPort != 443 && *ippermission.IpProtocol == "tcp" {
					for _, iprange := range ippermission.IpRanges {
						if *iprange.CidrIp == "0.0.0.0/0" || *iprange.CidrIp == "::/0" {
							fmt.Println("Group ID ", *group.GroupId, "with", *ippermission.IpProtocol, "/", *ippermission.FromPort, "allowing CidrIp", *iprange.CidrIp)

						}
					}

				}
			}
		}
	}

}

func exitErrorf(msg string, args ...interface{}) {
	fmt.Fprintf(os.Stderr, msg+"\n", args...)
	os.Exit(1)
}
