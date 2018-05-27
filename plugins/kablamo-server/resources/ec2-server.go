package resources

import (
	"github.com/KablamoOSS/kombustion/pluginParsers/resources"
	"github.com/KablamoOSS/kombustion/types"
	yaml "gopkg.in/yaml.v2"
)

type EC2Volume struct {
	Device   string `yaml:"Device"`
	VolumeId string `yaml:"VolumeId"`
}

//NetworkVPCConfig Main Object and construct
type EC2ServerConfig struct {
	Properties struct {
		AdditionalInfo                    string      `yaml:"AdditionalInfo,omitempty" example:"admin"`
		Affinity                          []string    `yaml:"Affinity,omitempty" example:"admin"`
		AvailabilityZone                  []string    `yaml:"AvailabilityZone,omitempty" example:"admin"`
		BlockDeviceMappings               []string    `yaml:"BlockDeviceMappings,omitempty" example:"admin"`
		CreditSpecification               interface{} `yaml:"CreditSpecification,omitempty" example:"admin"`
		DisableApiTermination             []string    `yaml:"DisableApiTermination,omitempty" example:"false"`
		EbsOptimized                      []string    `yaml:"EbsOptimized,omitempty" example:"false"`
		ElasticGpuSpecifications          []string    `yaml:"ElasticGpuSpecifications,omitempty" example:"admin"`
		HostId                            []string    `yaml:"HostId,omitempty" example:"admin"`
		IamInstanceProfile                []string    `yaml:"IamInstanceProfile,omitempty" example:"admin"`
		ImageId                           []string    `yaml:"ImageId,omitempty" example:"ami-79fd7eee"`
		InstanceInitiatedShutdownBehavior []string    `yaml:"InstanceInitiatedShutdownBehavior,omitempty" example:"admin"`
		InstanceType                      []string    `yaml:"InstanceType,omitempty" example:"admin"`
		Ipv6AddressCount                  []string    `yaml:"Ipv6AddressCount,omitempty" example:"admin"`
		Ipv6Addresses                     []string    `yaml:"Ipv6Addresses,omitempty" example:"admin"`
		KernelId                          []string    `yaml:"KernelId,omitempty" example:"admin"`
		KeyName                           []string    `yaml:"KeyName,omitempty" example:"testkey"`
		Monitoring                        []string    `yaml:"Monitoring,omitempty" example:"admin"`
		NetworkInterfaces                 []string    `yaml:"NetworkInterfaces,omitempty" example:"admin"`
		PlacementGroupName                []string    `yaml:"PlacementGroupName,omitempty" example:"admin"`
		PrivateIpAddress                  []string    `yaml:"PrivateIpAddress,omitempty" example:"192.168.1.1"`
		RamdiskId                         []string    `yaml:"RamdiskId,omitempty" example:"admin"`
		SecurityGroupIds                  []string    `yaml:"SecurityGroupIds,omitempty" example:"admin"`
		SourceDestCheck                   []string    `yaml:"SourceDestCheck,omitempty" example:"true"`
		SsmAssociations                   []string    `yaml:"SsmAssociations,omitempty" example:"admin"`
		SubnetId                          []string    `yaml:"SubnetId,omitempty" example:"admin"`
		Tags                              []string    `yaml:"Tags,omitempty" example:"admin"`
		Tenancy                           string      `yaml:"Tenancy,omitempty" example:"admin"`
		UserData                          []string    `yaml:"UserData,omitempty" example:"admin"`
		PreUserData                       []string    `yaml:"PreUserData,omitempty" example:"admin"`
		UserDataTemplate                  string      `yaml:"UserDataTemplate,omitempty" example:"admin"`
		PostUserData                      []string    `yaml:"PostUserData,omitempty" example:"admin"`
		Volumes                           []EC2Volume `yaml:"Volumes,omitempty" example:"admin"`
	} `yaml:"Properties"`
}

//ParseNetworkVPC parser builder.
func ParseEC2Server(name string, data string) (cf types.ValueMap, err error) {
	// Parse the config data
	var config EC2ServerConfig
	if err = yaml.Unmarshal([]byte(data), &config); err != nil {
		return
	}

	// validate the config
	config.Validate()

	// create a group of objects (each to be validated)
	cf = make(types.ValueMap)

	var serverData = map[string][]string{
		"Ubuntu1604 ": {
			"- '#!/bin/bash -v \n'",
			"{ 'Fn::Join': ['',['- sudo timedatectl set-timezone ',{'Ref': 'Timezone'},'\n']]},",
			"{ 'Fn::Join': ['',['- sudo hostnamectl set-hostname ',{'Ref': 'Hostname'},'\n']]},",
			"{ 'Fn::Join': ['',['- sudo export fqdn=',{'Ref': 'FQDN'},'\n']]},",
			"- echo $(hostname) >/home/ubuntu/isHostNameOk.log \n",
			"- echo 127.0.0.1 $(hostname)$(fqdn) $(hostname) >> /etc/hosts \n",
			"- apt-get update >/home/ubuntu/isUpdateOk.log \n",
			"- apt-get install -qy curl \n",
		},
		"Puppet": {
			"- '#!/bin/bash -v \n'",
			"{ 'Fn::Join': ['',['- sudo timedatectl set-timezone ',{'Ref': 'Timezone'},'\n']]},",
			"{ 'Fn::Join': ['',['- sudo hostnamectl set-hostname ',{'Ref': 'Hostname'},'\n']]},",
			"{ 'Fn::Join': ['',['- sudo export fqdn=',{'Ref': 'FQDN'},'\n']]},",
			"- echo $(hostname) >/home/ubuntu/isHostNameOk.log \n",
			"- echo 127.0.0.1 $(hostname)$fqdn $(hostname) >> /etc/hosts \n",
			"- apt-get update >/home/ubuntu/isUpdateOk.log \n",
			"- apt-get install -qy curl \n",
			"- export PE_CERTNAME=$(hostname)$(fqdn) \n",
			"{ 'Fn::Join': ['',['- export PE_ENV=',{'Ref': 'Puppet_Environment'},'\n']]},",
			"{ 'Fn::Join': ['',['- curl -k https://',{'Ref': 'Puppet_Master'},':8140/packages/current/install.bash | sudo bash -s main:ca_server=',{'Ref': 'Puppet_Master'},' agent:server=', {'Ref': 'Puppet_Master'},' agent:certname=${PE_CERTNAME} > /home/ubuntu/puppet-install.txt \n']]},",
			"- puppet config set environment $PE_ENV --section agent \n",
		},
	}

	userdataConfig := config.Properties.UserData

	if len(config.Properties.UserDataTemplate) < 0 {
		if len(config.Properties.PreUserData) < 0 {
			userdataConfig = userdataConfig
		}
		userdataConfig = serverData[config.Properties.UserDataTemplate]
	}

	cf[name] = resources.NewEC2Instance(
		resources.EC2InstanceProperties{
			AdditionalInfo:           config.Properties.AdditionalInfo,
			Affinity:                 config.Properties.Affinity,
			AvailabilityZone:         config.Properties.AvailabilityZone,
			BlockDeviceMappings:      config.Properties.BlockDeviceMappings,
			DisableApiTermination:    config.Properties.DisableApiTermination,
			EbsOptimized:             config.Properties.EbsOptimized,
			ElasticGpuSpecifications: config.Properties.ElasticGpuSpecifications,
			HostId:             config.Properties.HostId,
			IamInstanceProfile: config.Properties.IamInstanceProfile,
			ImageId:            config.Properties.ImageId,
			InstanceInitiatedShutdownBehavior: config.Properties.InstanceInitiatedShutdownBehavior,
			InstanceType:                      config.Properties.InstanceType,
			Ipv6AddressCount:                  config.Properties.Ipv6AddressCount,
			Ipv6Addresses:                     config.Properties.Ipv6Addresses,
			KernelId:                          config.Properties.KernelId,
			KeyName:                           config.Properties.KeyName,
			Monitoring:                        config.Properties.Monitoring,
			NetworkInterfaces:                 config.Properties.NetworkInterfaces,
			PlacementGroupName:                config.Properties.PlacementGroupName,
			PrivateIpAddress:                  config.Properties.PrivateIpAddress,
			RamdiskId:                         config.Properties.RamdiskId,
			SecurityGroupIds:                  config.Properties.SecurityGroupIds,
			SourceDestCheck:                   config.Properties.SourceDestCheck,
			SsmAssociations:                   config.Properties.SsmAssociations,
			SubnetId:                          config.Properties.SubnetId,
			Tags:                              config.Properties.Tags,
			Tenancy:                           config.Properties.Tenancy,
			UserData:                          userdataConfig,
			Volumes:                           config.Properties.Volumes,
		},
	)

	/* 	for k, resource := range cf {
	   		if errs := resource.Validate(); len(errs) > 0 {
	   			for _, err = range errs {
	   				log.Println("WARNING: KablamoNetworkConfig - ", err)
	   			}
	   			return
	   		}
	   	}
		   return */
	return
}

// Validate - input Config validation
func (this EC2ServerConfig) Validate() {

}
