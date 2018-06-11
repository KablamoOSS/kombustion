package outputs

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
)

// ParseRoute53RecordSetGroup Documentation http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-route53-recordsetgroup.html
func ParseRoute53RecordSetGroup(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {

	cf = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-Route53RecordSetGroup-" + name,
				},
			},
		},
	}

	return
}
