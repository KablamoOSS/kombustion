// +build plugin

package resources

import (
	"github.com/KablamoOSS/kombustion/pluginParsers/resources"
	"github.com/KablamoOSS/kombustion/types"
	yaml "gopkg.in/yaml.v2"
)

type WordpressConfig struct {
	Properties struct {
		InstanceType *string       `yaml:"InstanceType,omitempty"`
		Subnets      []interface{} `yaml:"Subnets"`
	} `yaml:"Properties"`
}

func ParseWordpress(name string, data string) (cf types.ValueMap, err error) {
	// Parse the config data
	var config WordpressConfig
	if err = yaml.Unmarshal([]byte(data), &config); err != nil {
		return
	}

	// validate the config
	config.Validate()

	// create a group of objects (each to be validated)
	cf = make(types.ValueMap)

	// defaults
	// instanceType := "t2.nano"
	// if config.Properties.InstanceType != nil {
	// instanceType = *config.Properties.InstanceType
	// }

	cf[name+"WebServerSecurityGroup"] = resources.NewEC2SecurityGroup(
		resources.EC2SecurityGroupProperties{
			GroupDescription: "Enable HTTP access via port 80 locked down to the load balancer + SSH access",
			SecurityGroupIngress: []interface{}{
				map[string]interface{}{"IpProtocol": "tcp"},
				map[string]interface{}{"FromPort": "80"},
				map[string]interface{}{"ToPort": "80"},
				map[interface{}]interface{}{"SourceSecurityGroupId": map[interface{}]interface{}{"Fn::Select": []interface{}{
					"0",
					map[interface{}]interface{}{"Fn::GetAtt": []interface{}{
						"ApplicationLoadBalancer",
						"SecurityGroups",
					}},
				}}},
				map[string]interface{}{"IpProtocol": "tcp"},
				map[string]interface{}{"FromPort": "22"},
				map[string]interface{}{"ToPort": "22"},
				map[interface{}]interface{}{"CidrIp": map[string]interface{}{"Ref": "SSHLocation"}},
			},
			VpcId: map[string]interface{}{"Ref": "VpcId"},
		},
	)

	cf[name+"WebServerGroup"] = resources.NewAutoScalingAutoScalingGroup(
		resources.AutoScalingAutoScalingGroupProperties{
			MaxSize:         "5",
			DesiredCapacity: map[string]interface{}{"Ref": "WebServerCapacity"},
			TargetGroupARNs: []interface{}{
				map[string]interface{}{"Ref": "ALBTargetGroup"},
			},
			VPCZoneIdentifier:       map[string]interface{}{"Ref": "Subnets"},
			LaunchConfigurationName: map[string]interface{}{"Ref": "LaunchConfig"},
			MinSize:                 "1",
		},
	)

	cf[name+"LaunchConfig"] = resources.NewAutoScalingLaunchConfiguration(
		resources.AutoScalingLaunchConfigurationProperties{
			InstanceType: map[string]interface{}{"Ref": "InstanceType"},
			SecurityGroups: []interface{}{
				map[string]interface{}{"Ref": "WebServerSecurityGroup"},
			},
			KeyName: map[string]interface{}{"Ref": "KeyName"},
			UserData: map[interface{}]interface{}{"Fn::Base64": map[interface{}]interface{}{"Fn::Join": []interface{}{
				"",
				[]interface{}{
					"#!/bin/bash -xe\n",
					"yum update -y aws-cfn-bootstrap\n",
					"/opt/aws/bin/cfn-init -v ",
					"         --stack ",
					map[string]interface{}{"Ref": "AWS::StackName"},
					"         --resource LaunchConfig ",
					"         --configsets wordpress_install ",
					"         --region ",
					map[string]interface{}{"Ref": "AWS::Region"},
					"\n",
					"/opt/aws/bin/cfn-signal -e $? ",
					"         --stack ",
					map[string]interface{}{"Ref": "AWS::StackName"},
					"         --resource WebServerGroup ",
					"         --region ",
					map[string]interface{}{"Ref": "AWS::Region"},
					"\n",
				}}}},
			ImageId: map[interface{}]interface{}{"Fn::FindInMap": []interface{}{
				"AWSRegionArch2AMI",
				map[string]interface{}{"Ref": "AWS::Region"},
				map[interface{}]interface{}{"Fn::FindInMap": []interface{}{
					"AWSInstanceType2Arch",
					map[string]interface{}{"Ref": "InstanceType"},
					"Arch",
				}},
			}},
		},
	)

	cf[name+"DBEC2SecurityGroup"] = resources.NewEC2SecurityGroup(
		resources.EC2SecurityGroupProperties{
			VpcId:            map[string]interface{}{"Ref": "VpcId"},
			GroupDescription: "Open database for access",
			SecurityGroupIngress: []interface{}{
				map[string]interface{}{"FromPort": "3306"},
				map[string]interface{}{"ToPort": "3306"},
				map[interface{}]interface{}{"SourceSecurityGroupId": map[string]interface{}{"Ref": "WebServerSecurityGroup"}},
				map[string]interface{}{"IpProtocol": "tcp"},
			},
		},
	)

	cf[name+"DBInstance"] = resources.NewRDSDBInstance(
		resources.RDSDBInstanceProperties{
			DBInstanceClass:  map[string]interface{}{"Ref": "DBClass"},
			AllocatedStorage: map[string]interface{}{"Ref": "DBAllocatedStorage"},
			VPCSecurityGroups: []interface{}{
				map[interface{}]interface{}{"Fn::GetAtt": []interface{}{
					"DBEC2SecurityGroup",
					"GroupId",
				}},
			},
			DBName:             map[string]interface{}{"Ref": "DBName"},
			Engine:             "MySQL",
			MultiAZ:            map[string]interface{}{"Ref": "MultiAZDatabase"},
			MasterUsername:     map[string]interface{}{"Ref": "DBUser"},
			MasterUserPassword: map[string]interface{}{"Ref": "DBPassword"},
		},
	)

	cf[name+"ApplicationLoadBalancer"] = resources.NewElasticLoadBalancingV2LoadBalancer(
		resources.ElasticLoadBalancingV2LoadBalancerProperties{
			Subnets: map[string]interface{}{"Ref": "Subnets"},
		},
	)

	cf[name+"ALBListener"] = resources.NewElasticLoadBalancingV2Listener(
		resources.ElasticLoadBalancingV2ListenerProperties{
			DefaultActions: []interface{}{
				map[string]interface{}{"Type": "forward"},
				map[interface{}]interface{}{"TargetGroupArn": map[string]interface{}{"Ref": "ALBTargetGroup"}},
			},
			LoadBalancerArn: map[string]interface{}{"Ref": "ApplicationLoadBalancer"},
			Port:            "80",
			Protocol:        "HTTP",
		},
	)

	cf[name+"ALBTargetGroup"] = resources.NewElasticLoadBalancingV2TargetGroup(
		resources.ElasticLoadBalancingV2TargetGroupProperties{
			HealthyThresholdCount:   "2",
			UnhealthyThresholdCount: "5",
			TargetGroupAttributes: []interface{}{
				map[string]interface{}{"Key": "stickiness.enabled"},
				map[string]interface{}{"Value": "true"},
				map[string]interface{}{"Key": "stickiness.type"},
				map[string]interface{}{"Value": "lb_cookie"},
				map[string]interface{}{"Key": "stickiness.lb_cookie.duration_seconds"},
				map[string]interface{}{"Value": "30"},
			},
			HealthCheckTimeoutSeconds:  "5",
			HealthCheckIntervalSeconds: "10",
			Port:            "80",
			Protocol:        "HTTP",
			VpcId:           map[string]interface{}{"Ref": "VpcId"},
			HealthCheckPath: "/wordpress/wp-admin/install.php",
		},
	)

	return
}

// Validate - input Config validation
func (this WordpressConfig) Validate() {

}
