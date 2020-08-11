package service

import (
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/spf13/viper"
)

//EC2Service service for various functions provided by app for EC2
type EC2Service struct {
	ec2 *ec2.EC2
}

//NewEC2Service create new EC2Service instance
func NewEC2Service() *EC2Service {
	awsSession := viper.Get("awsSession").(*session.Session)
	ec2Client := ec2.New(awsSession)
	ec2Service := EC2Service{ec2: ec2Client}
	return &ec2Service
}

//VerifySecurityGroups get security group description
func (ec2Service EC2Service) VerifySecurityGroups() {
	input := &ec2.DescribeSecurityGroupsInput{
		GroupIds: getSecurityGrps(),
	}
	result, err := ec2Service.ec2.DescribeSecurityGroups(input)
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
	formattedDescription(result.SecurityGroups)
}

func formattedDescription(securityGrps []*ec2.SecurityGroup) {
	fmt.Println("Description of security groups in configuration file")
	for _, securityGrp := range securityGrps {
		fmt.Printf("[%s]:[%s]:[%s]\n", *securityGrp.Description, *securityGrp.GroupId, *securityGrp.GroupName)
		fmt.Println(">>>>>>>>>>>>>>>>>>>>> Open Ports >>>>>>>>>>>>>>>>>>>>>>>>>>>")
		fmt.Printf("<<<<\t[From]\t\t[To]\t\t[Protocol]\t>>>>\n")
		for _, ipp := range securityGrp.IpPermissions {
			fmt.Printf("<<<<\t%d\t\t%d\t\t%s\t\t>>>>\n", *ipp.FromPort, *ipp.ToPort, *ipp.IpProtocol)
		}

		fmt.Println("<<<<<<<<<<<<<<<<<<<<< Open Ports <<<<<<<<<<<<<<<<<<<<<<<<<<<")
		fmt.Println()
	}
}
func getSecurityGrps() []*string {
	securityGrpIds := viper.Get("aws.ec2.security.groups").([]interface{})
	awsSecurityGrpIds := make([]*string, len(securityGrpIds))
	for _, x := range securityGrpIds {
		awsSecurityGrpIds = append(awsSecurityGrpIds, aws.String(x.(string)))
	}
	return awsSecurityGrpIds
}

//UpdateIPWith this method filters IP Permissions in configured security groups
//by :ruleDescription and then update each IP Permission with :newIP
func (ec2Service EC2Service) UpdateIPWith(newIP, ruleDescription string) {
	log.Printf("updating ruleDescription %s with ip %s \n", ruleDescription, newIP)
}
