// +build plugin

package resources

import (
	"log"
	"strings"

	"github.com/KablamoOSS/kombustion/pluginParsers/properties"
	"github.com/KablamoOSS/kombustion/pluginParsers/resources"
	"github.com/KablamoOSS/kombustion/types"
	yaml "gopkg.in/yaml.v2"
)

//Subnet object
type Subnet struct {
	CIDR       string `yaml:"CIDR"`
	AZ         string `yaml:"AZ"`
	NetACL     string `yaml:"NetACL"`
	RouteTable string `yaml:"RouteTable"`
}

//Routetable an array of Route(s)
type Routetable struct {
	routes []Route `yaml:"routes"`
}

//Route Object
type Route struct {
	routename string `yaml:"routename"`
	routecidr string `yaml:"routecidr"`
	routegw   string `yaml:"routegw"`
}

//NetworkVPCConfig Main Object and construct
type NetworkVPCConfig struct {
	Properties struct {
		CIDR *string `yaml:"CIDR"`
		DHCP struct {
			name           string `yaml:"name"`
			dnsservers     string `yaml:"dnsservers"`
			ntpservers     string `yaml:"ntpservers,omitempty"`
			ntbtype        string `yaml:"ntbtype,omitempty"`
			domainname     string `yaml:"domainname,omitempty"`
			netbiosservers string `yaml:"netbiosservers,omitempty"`
		} `yaml:"DHCP"`
		Details struct {
			VPCName string `yaml:"VPCName"`
			VPCDesc string `yaml:"VPCDesc"`
			Region  string `yaml:"Region"`
		} `yaml:"Details"`
		Subnets     map[string]Subnet      `yaml:"Subnets,omitempty"`
		RouteTables map[string]Routetable  `yaml:"RouteTables,omitempty"`
		NetworkACLs map[string]interface{} `yaml:"NetworkACLs,omitempty"`
		Tags        interface{}            `yaml:"Tags"`
	} `yaml:"Properties"`
}

//ParseNetworkVPC parser builder.
func ParseNetworkVPC(name string, data string) (cf types.ValueMap, err error) {
	// Parse the config data
	var config NetworkVPCConfig
	if err = yaml.Unmarshal([]byte(data), &config); err != nil {
		return
	}

	// validate the config
	config.Validate()

	// create a group of objects (each to be validated)
	cf = make(types.ValueMap)

	cf[config.Properties.Details.VPCName] = resources.NewEC2VPC(
		resources.EC2VPCProperties{
			CidrBlock:          config.Properties.CIDR,
			EnableDnsHostnames: true,
			EnableDnsSupport:   true,
			InstanceTenancy:    true,
		},
	)

	cf[config.Properties.DHCP.name] = resources.NewEC2DHCPOptions(
		resources.EC2DHCPOptionsProperties{
			DomainName:         config.Properties.DHCP.domainname,
			NetbiosNodeType:    config.Properties.DHCP.ntbtype,
			DomainNameServers:  config.Properties.DHCP.dnsservers,
			NetbiosNameServers: config.Properties.DHCP.netbiosservers,
			NtpServers:         config.Properties.DHCP.ntpservers,
			Tags:               map[string]string{"Name": config.Properties.DHCP.name},
		},
	)

	cf[config.Properties.DHCP.name+"VPCDHCPOptionsAssociation"] = resources.NewEC2VPCDHCPOptionsAssociation(
		resources.EC2VPCDHCPOptionsAssociationProperties{
			DhcpOptionsId: map[string]interface{}{"Ref": config.Properties.DHCP.name},
			VpcId:         map[string]interface{}{"Ref": config.Properties.Details.VPCName},
		},
	)

	cf["InternetGateway"] = resources.NewEC2InternetGateway(
		resources.EC2InternetGatewayProperties{
			Tags: map[string]interface{}{"Name": "IGW"},
		},
	)

	cf["InternetGatewayVPCGatewayAttachment"] = resources.NewEC2VPCGatewayAttachment(
		resources.EC2VPCGatewayAttachmentProperties{
			InternetGatewayId: map[string]interface{}{"Ref": "InternetGateway"},
			VpcId:             map[string]interface{}{"Ref": config.Properties.Details.VPCName},
		},
	)

	cf["VPNGatewayVPCGatewayAttachment"] = resources.NewEC2VPCGatewayAttachment(
		resources.EC2VPCGatewayAttachmentProperties{
			VpnGatewayId: map[string]interface{}{"Ref": "VGW"},
			VpcId:        map[string]interface{}{"Ref": config.Properties.Details.VPCName},
		},
	)

	for routetable, settings := range config.Properties.RouteTables {
		cf[routetable] = resources.NewEC2RouteTable(
			resources.EC2RouteTableProperties{
				VpcId: config.Properties.Details.VPCName,
				Tags:  map[string]string{"Name": routetable},
			},
		)

		for route, routedetail := range settings.routes {
			if route > 0 {
				cf[routetable+routedetail.routename] = resources.NewEC2Route(
					resources.EC2RouteProperties{
						DestinationCidrBlock: routedetail.routecidr,
						GatewayId:            map[string]string{"Ref": routedetail.routegw},
					},
				)
			}
		}

		cf[routetable+"RoutePropagation"] = resources.NewEC2VPNGatewayRoutePropagation(
			resources.EC2VPNGatewayRoutePropagationProperties{
				RouteTableIds: map[string]string{"Ref": routetable},
				VpnGatewayId:  map[string]string{"Ref": "VGW"},
			},
			"VPNGatewayVPCGatewayAttachment",
		)
	}

	for networkacl, settings := range config.Properties.NetworkACLs {
		cf[networkacl] = resources.NewEC2NetworkAcl(
			resources.EC2NetworkAclProperties{
				VpcId: map[string]interface{}{"Ref": config.Properties.Details.VPCName},
				Tags:  map[string]interface{}{"Name": networkacl},
			},
		)

		for aclentry, acl := range settings.(map[interface{}]interface{}) {
			ports := properties.NetworkAclEntry_PortRange{
				From: strings.Split(acl.(string), ",")[5],
				To:   strings.Split(acl.(string), ",")[6],
			}
			cf[aclentry.(string)] = resources.NewEC2NetworkAclEntry(
				resources.EC2NetworkAclEntryProperties{
					NetworkAclId: map[string]interface{}{"Ref": networkacl},
					RuleNumber:   strings.Split(acl.(string), ",")[0],
					Protocol:     strings.Split(acl.(string), ",")[1],
					RuleAction:   strings.Split(acl.(string), ",")[2],
					Egress:       strings.Split(acl.(string), ",")[3],
					CidrBlock:    strings.Split(acl.(string), ",")[4],
					PortRange:    &ports,
				},
			)
		}
	}

	for subnet, settings := range config.Properties.Subnets {
		cf[subnet] = resources.NewEC2Subnet(
			resources.EC2SubnetProperties{
				VpcId:            map[string]interface{}{"Ref": config.Properties.Details.VPCName},
				CidrBlock:        settings.CIDR,
				AvailabilityZone: map[interface{}]interface{}{"Fn::Select": []interface{}{settings.AZ, "Fn::GetAZs"}},
				Tags:             map[string]string{"Name": subnet},
			},
		)

		cf[subnet+"SubnetNetworkAclAssociation"] = resources.NewEC2SubnetNetworkAclAssociation(
			resources.EC2SubnetNetworkAclAssociationProperties{
				NetworkAclId: map[string]interface{}{"Ref": settings.NetACL},
				SubnetId:     map[string]interface{}{"Ref": subnet},
			},
		)

		cf[subnet+"SubnetRouteTableAssociation"] = resources.NewEC2SubnetRouteTableAssociation(
			resources.EC2SubnetRouteTableAssociationProperties{
				RouteTableId: map[string]interface{}{"Ref": settings.RouteTable},
				SubnetId:     map[string]interface{}{"Ref": subnet},
			},
		)
	}

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
func (this NetworkVPCConfig) Validate() {
	if this.Properties.CIDR == nil {
		log.Println("WARNING: KablamoNetworkConfig - Missing required field 'CIDR'")
	}
}
