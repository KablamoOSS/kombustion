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

//NATGateway object
type NATGateway struct {
	Subnet     string `yaml:"Subnet"`
	Routetable string `yaml:"Routetable"`
}

//Routetable an array of Route(s)
type Routetable struct {
	routes []Route `yaml:"routes"`
}

//Route Object
type Route struct {
	RouteName string `yaml:"RouteName"`
	RouteCIDR string `yaml:"RouteCIDR"`
	RouteGW   string `yaml:"RouteGW"`
}

type Tag struct {
	Key   string `yaml:"Key"`
	Value string `yaml:"Value"`
}

//NetworkVPCConfig Main Object and construct
type NetworkVPCConfig struct {
	Properties struct {
		CIDR *string `yaml:"CIDR"`
		DHCP struct {
			Name           string `yaml:"Name"`
			DNSServers     string `yaml:"DNSServers"`
			NTPServers     string `yaml:"NTPServers,omitempty"`
			NTBType        string `yaml:"NTBType,omitempty"`
			Domainname     string `yaml:"Domainname,omitempty"`
			Netbiosservers string `yaml:"Netbiosservers,omitempty"`
		} `yaml:"DHCP"`
		Details struct {
			VPCName string `yaml:"VPCName"`
			VPCDesc string `yaml:"VPCDesc"`
			Region  string `yaml:"Region"`
		} `yaml:"Details"`
		Subnets     map[string]Subnet      `yaml:"Subnets,omitempty"`
		NatGateways map[string]NATGateway  `yaml:"NATGateways,omitempty"`
		RouteTables map[string][]Route     `yaml:"RouteTables,omitempty"`
		NetworkACLs map[string]interface{} `yaml:"NetworkACLs,omitempty"`
		Tags        interface{}            `yaml:"Tags"`
	} `yaml:"Properties"`
}

func splitStrArray(asset string) []string {
	if len(asset) > 0 {
		strArray := strings.Split(asset, ",")
		return strArray
	} else {
		return nil
	}
}

func genMap(asset map[string]string) []map[string]string {
	arrayMap := make([]map[string]string, 0)
	arrayMap = append(arrayMap, asset)
	return arrayMap
}

func genTags(Tags map[string]string) []Tag {
	arrayTags := make([]Tag, 0)
	for k, v := range Tags {
		arrayTags = append(arrayTags, Tag{Key: k, Value: v})
	}
	return arrayTags
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
			InstanceTenancy:    "default",
			Tags:               genTags(map[string]string{"Name": config.Properties.Details.VPCName}),
		},
	)

	cf[config.Properties.DHCP.Name] = resources.NewEC2DHCPOptions(
		resources.EC2DHCPOptionsProperties{
			DomainName:         config.Properties.DHCP.Name,
			NetbiosNodeType:    config.Properties.DHCP.NTBType,
			DomainNameServers:  splitStrArray(config.Properties.DHCP.DNSServers),
			NetbiosNameServers: splitStrArray(config.Properties.DHCP.Netbiosservers),
			NtpServers:         splitStrArray(config.Properties.DHCP.NTPServers),
			Tags:               genTags(map[string]string{"Name": config.Properties.DHCP.Name}),
		},
	)

	cf[config.Properties.DHCP.Name+"VPCDHCPOptionsAssociation"] = resources.NewEC2VPCDHCPOptionsAssociation(
		resources.EC2VPCDHCPOptionsAssociationProperties{
			DhcpOptionsId: map[string]interface{}{"Ref": config.Properties.DHCP.Name},
			VpcId:         map[string]interface{}{"Ref": config.Properties.Details.VPCName},
		},
	)

	cf["InternetGateway"] = resources.NewEC2InternetGateway(
		resources.EC2InternetGatewayProperties{
			Tags: genTags(map[string]string{"Name": "IGW"}),
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
				VpcId: map[string]interface{}{"Ref": config.Properties.Details.VPCName},
				Tags:  genTags(map[string]string{"Name": routetable}),
			},
		)

		for _, routeinfo := range settings {
			cf[routeinfo.RouteName] = resources.NewEC2Route(
				resources.EC2RouteProperties{
					DestinationCidrBlock: routeinfo.RouteCIDR,
					GatewayId:            map[string]string{"Ref": routeinfo.RouteGW},
					RouteTableId:         map[string]string{"Ref": routetable},
				},
			)
		}

		cf[routetable+"RoutePropagation"] = resources.NewEC2VPNGatewayRoutePropagation(
			resources.EC2VPNGatewayRoutePropagationProperties{
				RouteTableIds: genMap(map[string]string{"Ref": routetable}),
				VpnGatewayId:  map[string]string{"Ref": "VGW"},
			},
			"VPNGatewayVPCGatewayAttachment",
		)
	}

	for networkacl, settings := range config.Properties.NetworkACLs {
		cf[networkacl] = resources.NewEC2NetworkAcl(
			resources.EC2NetworkAclProperties{
				VpcId: map[string]interface{}{"Ref": config.Properties.Details.VPCName},
				Tags:  genTags(map[string]string{"Name": networkacl}),
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
				AvailabilityZone: map[interface{}]interface{}{"Fn::Select": []interface{}{settings.AZ, map[string]string{"Fn::GetAZs": ""}}},
				Tags:             genTags(map[string]string{"Name": subnet}),
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

	for natgw, settings := range config.Properties.NatGateways {
		cf["EIP"+natgw] = resources.NewEC2EIP(
			resources.EC2EIPProperties{
				Domain: "vpc",
			},
		)

		cf[natgw] = resources.NewEC2NatGateway(
			resources.EC2NatGatewayProperties{
				AllocationId: map[string]interface{}{"Fn::GetAtt": []string{"EIP" + natgw, "AllocationId"}},
				SubnetId:     map[string]interface{}{"Ref": settings.Subnet},
				Tags:         genTags(map[string]string{"Name": natgw}),
			},
		)

		cf[natgw+"Route"] = resources.NewEC2Route(
			resources.EC2RouteProperties{
				DestinationCidrBlock: "0.0.0.0/0",
				RouteTableId:         map[string]string{"Ref": settings.Routetable},
				NatGatewayId:         map[string]string{"Ref": natgw},
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
