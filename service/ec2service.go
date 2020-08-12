package service

import (
	"fmt"
	"log"
	"strings"

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
	securityGrps := ec2Service.fetchSecurityGrps()
	formattedDescription(securityGrps)
}
func (ec2Service EC2Service) fetchSecurityGrps() []*ec2.SecurityGroup {
	input := &ec2.DescribeSecurityGroupsInput{
		GroupIds: getSecurityGrpsIds(),
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
	}
	return result.SecurityGroups
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
	}
}
func getSecurityGrpsIds() []*string {
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
	securityGrps := ec2Service.fetchSecurityGrps()
	formattedDescription(securityGrps)
	log.Printf("updating ruleDescription %s with ip %s \n", ruleDescription, newIP)
	for _, securityGrp := range securityGrps {
		for _, ipp := range securityGrp.IpPermissions {
			if hasRangeWith(ruleDescription, ipp) {
				ec2Service.revokePermissionForOldIP(ruleDescription, ipp, securityGrp)
			}
		}
	}
	securityGrps = ec2Service.fetchSecurityGrps()
	log.Printf("Revoked permissions for ruleDescription %s with ip %s \n", ruleDescription, newIP)
	formattedDescription(securityGrps)
	for _, securityGrp := range securityGrps {
		for _, ipp := range securityGrp.IpPermissions {
			ec2Service.createNewIPPermission(newIP, ruleDescription, ipp, *securityGrp.GroupId)
		}
	}
	securityGrps = ec2Service.fetchSecurityGrps()
	log.Printf("Added permissions for ruleDescription %s with ip %s \n", ruleDescription, newIP)
	formattedDescription(securityGrps)
}
func (ec2Service EC2Service) createNewIPPermission(newIP, ruleDescription string, oldIpp *ec2.IpPermission, securityGrpID string) {
	log.Printf("Creating new permissions in %s for for range from %d to %d \n", securityGrpID, *oldIpp.FromPort, *oldIpp.ToPort)
	ec2Service.addNewPermission(newIP, ruleDescription, oldIpp, securityGrpID)
	ec2Service.updateDescription(newIP, ruleDescription, oldIpp, securityGrpID)
	log.Printf("Created new permissions in %s for for range from %d to %d \n", securityGrpID, *oldIpp.FromPort, *oldIpp.ToPort)
}
func (ec2Service EC2Service) updateDescription(newIP, ruleDescription string, oldIpp *ec2.IpPermission, securityGrpID string) {
	ipRange := &ec2.IpRange{
		CidrIp:      toCidrIP(newIP),
		Description: &ruleDescription,
	}
	newIpp := &ec2.IpPermission{
		IpProtocol: oldIpp.IpProtocol,
		FromPort:   oldIpp.FromPort,
		ToPort:     oldIpp.ToPort,
		IpRanges:   append([]*ec2.IpRange{}, ipRange),
	}
	updateDescriptionInput := &ec2.UpdateSecurityGroupRuleDescriptionsIngressInput{
		GroupId:       &securityGrpID,
		IpPermissions: append([]*ec2.IpPermission{}, newIpp),
	}
	ec2Service.ec2.UpdateSecurityGroupRuleDescriptionsIngress(updateDescriptionInput)
	log.Printf("Updated rule description in %s for range from %d to %d \n", securityGrpID, *oldIpp.FromPort, *oldIpp.ToPort)
}
func (ec2Service EC2Service) addNewPermission(newIP, ruleDescription string, oldIpp *ec2.IpPermission, securityGrpID string) {
	input := &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId:    &securityGrpID,
		IpProtocol: oldIpp.IpProtocol,
		FromPort:   oldIpp.FromPort,
		ToPort:     oldIpp.ToPort,
		CidrIp:     toCidrIP(newIP),
	}
	ec2Service.ec2.AuthorizeSecurityGroupIngress(input)
	log.Printf("Added new permissions in %s for ruleDescription %s with ip %s \n", securityGrpID, ruleDescription, newIP)
}
func toCidrIP(newIP string) *string {
	cidrIP := newIP + "/32"
	return &cidrIP
}
func (ec2Service EC2Service) revokePermissionForOldIP(ruleDescription string, ipp *ec2.IpPermission, securityGrp *ec2.SecurityGroup) {
	for _, ipRange := range ipp.IpRanges {
		if strings.EqualFold(ruleDescription, *ipRange.Description) {
			input := &ec2.RevokeSecurityGroupIngressInput{
				GroupId:    securityGrp.GroupId,
				IpProtocol: ipp.IpProtocol,
				FromPort:   ipp.FromPort,
				ToPort:     ipp.ToPort,
				CidrIp:     ipRange.CidrIp,
			}
			ec2Service.ec2.RevokeSecurityGroupIngress(input)
		}
	}
}
func hasRangeWith(ruleDescription string, ipp *ec2.IpPermission) bool {
	for _, ipRange := range ipp.IpRanges {
		if strings.EqualFold(ruleDescription, *ipRange.Description) {
			return true
		}
	}
	return false
}
