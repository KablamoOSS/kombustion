package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type SSMPatchBaseline struct {
	Type       string                      `yaml:"Type"`
	Properties SSMPatchBaselineProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type SSMPatchBaselineProperties struct {
	ApprovedPatchesComplianceLevel interface{} `yaml:"ApprovedPatchesComplianceLevel,omitempty"`
	ApprovedPatchesEnableNonSecurity interface{} `yaml:"ApprovedPatchesEnableNonSecurity,omitempty"`
	Description interface{} `yaml:"Description,omitempty"`
	Name interface{} `yaml:"Name"`
	OperatingSystem interface{} `yaml:"OperatingSystem,omitempty"`
	ApprovalRules *properties.PatchBaseline_RuleGroup `yaml:"ApprovalRules,omitempty"`
	GlobalFilters *properties.PatchBaseline_PatchFilterGroup `yaml:"GlobalFilters,omitempty"`
	ApprovedPatches interface{} `yaml:"ApprovedPatches,omitempty"`
	PatchGroups interface{} `yaml:"PatchGroups,omitempty"`
	RejectedPatches interface{} `yaml:"RejectedPatches,omitempty"`
	Sources interface{} `yaml:"Sources,omitempty"`
}

func NewSSMPatchBaseline(properties SSMPatchBaselineProperties, deps ...interface{}) SSMPatchBaseline {
	return SSMPatchBaseline{
		Type:       "AWS::SSM::PatchBaseline",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseSSMPatchBaseline(name string, data string) (cf types.ValueMap, err error) {
	var resource SSMPatchBaseline
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SSMPatchBaseline - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource SSMPatchBaseline) Validate() []error {
	return resource.Properties.Validate()
}

func (resource SSMPatchBaselineProperties) Validate() []error {
	errs := []error{}
	if resource.Name == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Name'"))
	}
	return errs
}
