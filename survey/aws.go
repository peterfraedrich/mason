package survey

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"runtime"
	"strings"
	"time"
)

var c = &http.Client{Timeout: 10 * time.Millisecond}

func awsClient(route string) string {
	url := "http://169.254.169.254/latest/" + route
	resp, err := c.Get(url)
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(body)
}

//AWSData type
type AWSData struct {
	Ec2AmiID            string   `json:"ec2_ami_id,omitempty"`
	Ec2AvailabilityZone string   `json:"ec2_availability_zone,omitempty"`
	Ec2InstanceID       string   `json:"ec2_instance_id,omitempty"`
	Ec2InstanceType     string   `json:"ec2_instance_type,omitempty"`
	Ec2Profile          string   `json:"ec2_profile,omitempty"`
	Ec2PublicIP4        string   `json:"ec2_public_ip4,omitempty"`
	Ec2SecurityGroups   []string `json:"ec2_security_groups,omitempty"`
	Ec2IAMID            string   `json:"ec2_iam_id,omitempty"`
	Ec2IAMARN           string   `json:"ec2_iam_arn,omitempty"`
}

// GetAWS grabs meta-data from AWS instance
func GetAWS() (AWSData, error) {
	if runtime.GOOS == "darwin" {
		return AWSData{}, nil
	}
	awsResponse, err := c.Get("http://169.254.169.254/latest/")
	if err != nil {
		return AWSData{}, err
	}
	if awsResponse == nil || awsResponse.StatusCode != http.StatusOK {
		return AWSData{}, errors.New("aws get: unable to get metadata from AWS")
	}
	d := AWSData{}
	d.Ec2AmiID = awsClient("meta-data/ami-id")
	d.Ec2InstanceID = awsClient("meta-data/instance-id")
	d.Ec2InstanceType = awsClient("meta-data/instance-type")
	d.Ec2AvailabilityZone = awsClient("meta-data/placement/availability-zone")
	d.Ec2Profile = awsClient("meta-data/profile")
	d.Ec2PublicIP4 = awsClient("meta-data/public-ipv4")
	d.Ec2SecurityGroups = strings.Split(strings.TrimSpace(awsClient("meta-data/security-groups")), "\n")
	iamInfo := struct {
		Code               string
		LastUpdated        string
		InstanceProfileArn string
		InstanceProfileID  string
	}{}
	rawIAMInfo := awsClient("meta-data/iam/info")
	err = json.Unmarshal([]byte(rawIAMInfo), &iamInfo)
	if err != nil {
		return d, err
	}
	if iamInfo.Code != "Success" {
		return d, fmt.Errorf("aws get: got code %s when trying to read IAM Role info", iamInfo.Code)
	}
	d.Ec2IAMARN = iamInfo.InstanceProfileArn
	d.Ec2IAMID = iamInfo.InstanceProfileID
	return d, nil
}
