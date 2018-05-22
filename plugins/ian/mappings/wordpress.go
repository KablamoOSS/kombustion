// +build plugin

package mappings

import (
	"github.com/KablamoOSS/kombustion/types"
)

func ParseWordpress(name string, data string) (cf types.ValueMap, err error) {
	// create a group of objects (each to be validated)
	cf = make(types.ValueMap)

	cf["WordpressAWSRegionArch2AMI"] = map[string]map[string]string{
		"us-east-1": map[string]string{
			"PV64":  "ami-2a69aa47",
			"HVM64": "ami-6869aa05",
			"HVMG2": "ami-3353c649",
		},
		"us-west-2": map[string]string{
			"PV64":  "ami-7f77b31f",
			"HVM64": "ami-7172b611",
			"HVMG2": "ami-58ce1220",
		},
		"us-west-1": map[string]string{
			"PV64":  "ami-a2490dc2",
			"HVM64": "ami-31490d51",
			"HVMG2": "ami-62ad9502",
		},
		"eu-west-1": map[string]string{
			"PV64":  "ami-4cdd453f",
			"HVM64": "ami-f9dd458a",
			"HVMG2": "ami-41bc0a38",
		},
		"eu-west-2": map[string]string{
			"PV64":  "NOT_SUPPORTED",
			"HVM64": "ami-886369ec",
			"HVMG2": "NOT_SUPPORTED",
		},
		"eu-west-3": map[string]string{
			"PV64":  "NOT_SUPPORTED",
			"HVM64": "NOT_SUPPORTED",
			"HVMG2": "NOT_SUPPORTED",
		},
		"eu-central-1": map[string]string{
			"PV64":  "ami-6527cf0a",
			"HVM64": "ami-ea26ce85",
			"HVMG2": "ami-b50d8fda",
		},
		"ap-northeast-1": map[string]string{
			"PV64":  "ami-3e42b65f",
			"HVM64": "ami-374db956",
			"HVMG2": "ami-14e45872",
		},
		"ap-northeast-2": map[string]string{
			"PV64":  "NOT_SUPPORTED",
			"HVM64": "ami-2b408b45",
			"HVMG2": "NOT_SUPPORTED",
		},
		"ap-southeast-1": map[string]string{
			"PV64":  "ami-df9e4cbc",
			"HVM64": "ami-a59b49c6",
			"HVMG2": "ami-2a80d649",
		},
		"ap-southeast-2": map[string]string{
			"PV64":  "ami-63351d00",
			"HVM64": "ami-dc361ebf",
			"HVMG2": "ami-02c42e60",
		},
		"ap-south-1": map[string]string{
			"PV64":  "NOT_SUPPORTED",
			"HVM64": "ami-ffbdd790",
			"HVMG2": "ami-f6165899",
		},
		"us-east-2": map[string]string{
			"PV64":  "NOT_SUPPORTED",
			"HVM64": "ami-f6035893",
			"HVMG2": "NOT_SUPPORTED",
		},
		"ca-central-1": map[string]string{
			"PV64":  "NOT_SUPPORTED",
			"HVM64": "ami-730ebd17",
			"HVMG2": "NOT_SUPPORTED",
		},
		"sa-east-1": map[string]string{
			"PV64":  "ami-1ad34676",
			"HVM64": "ami-6dd04501",
			"HVMG2": "NOT_SUPPORTED",
		},
		"cn-north-1": map[string]string{
			"PV64":  "ami-77559f1a",
			"HVM64": "ami-8e6aa0e3",
			"HVMG2": "NOT_SUPPORTED",
		},
		"cn-northwest-1": map[string]string{
			"PV64":  "ami-80707be2",
			"HVM64": "ami-cb858fa9",
			"HVMG2": "NOT_SUPPORTED",
		},
	}

	return
}
