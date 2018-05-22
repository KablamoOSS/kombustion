// +build plugin

package resources

import (
	"strings"

	"github.com/KablamoOSS/kombustion/pluginParsers/resources"
	"github.com/KablamoOSS/kombustion/types"
	yaml "gopkg.in/yaml.v2"
)

//NetworkVPCConfig Main Object and construct
type NetworkSecurityGroupsConfig struct {
	Properties struct {
		Description string   `yaml:"Description"`
		Ingress     []string `yaml:"Ingress"`
		Egress      []string `yaml:"Egress"`
	} `yaml:"Properties"`
}

func splitSecGroupObject(data string, sectype string) (secobj interface{}) {
	secentry := strings.Split(data, ",")
	switch entrysize := len(secentry); entrysize {
	case 3:
		secentry = append(secentry, "0.0.0.0/0")
		secentry = append(secentry, "")
	case 4:
		secentry = append(secentry, "")
	}
	if sectype == "ingress" {
		secobj = resources.EC2SecurityGroupIngressProperties{
			CidrIp:      map[string]interface{}{"Ref": secentry[3]},
			FromPort:    secentry[1],
			ToPort:      secentry[2],
			IpProtocol:  secentry[0],
			Description: secentry[4],
		}
	} else {
		secobj = resources.EC2SecurityGroupEgressProperties{
			CidrIp:      map[string]interface{}{"Ref": secentry[3]},
			FromPort:    secentry[1],
			ToPort:      secentry[2],
			IpProtocol:  secentry[0],
			Description: secentry[4],
		}
	}
	return secobj
}

//ParseNetworkVPC parser builder.
func ParseNetworkSecurityGroups(name string, data string) (cf types.ValueMap, err error) {
	// Parse the config data
	var config NetworkSecurityGroupsConfig
	if err = yaml.Unmarshal([]byte(data), &config); err != nil {
		return
	}

	// validate the config
	config.Validate()

	// create a group of objects (each to be validated)
	cf = make(types.ValueMap)

	var ingress = []resources.EC2SecurityGroupIngressProperties{}
	for _, obj := range config.Properties.Ingress {
		ingress = append(
			ingress, splitSecGroupObject(obj, "ingress").(resources.EC2SecurityGroupIngressProperties),
		)
	}

	var egress = []resources.EC2SecurityGroupEgressProperties{}
	for _, obj := range config.Properties.Egress {
		egress = append(
			egress, splitSecGroupObject(obj, "egress").(resources.EC2SecurityGroupEgressProperties),
		)
	}

	cf[name] = resources.NewEC2SecurityGroup(
		resources.EC2SecurityGroupProperties{
			GroupDescription:     config.Properties.Description,
			SecurityGroupIngress: ingress,
			SecurityGroupEgress:  egress,
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
func (this NetworkSecurityGroupsConfig) Validate() {

}
