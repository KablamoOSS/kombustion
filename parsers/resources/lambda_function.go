package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
	"fmt"
	"github.com/KablamoOSS/kombustion/parsers/properties"
)

type LambdaFunction struct {
	Type       string                      `yaml:"Type"`
	Properties LambdaFunctionProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type LambdaFunctionProperties struct {
	Description interface{} `yaml:"Description,omitempty"`
	FunctionName interface{} `yaml:"FunctionName,omitempty"`
	Handler interface{} `yaml:"Handler"`
	KmsKeyArn interface{} `yaml:"KmsKeyArn,omitempty"`
	MemorySize interface{} `yaml:"MemorySize,omitempty"`
	ReservedConcurrentExecutions interface{} `yaml:"ReservedConcurrentExecutions,omitempty"`
	Role interface{} `yaml:"Role"`
	Runtime interface{} `yaml:"Runtime"`
	Timeout interface{} `yaml:"Timeout,omitempty"`
	VpcConfig *properties.Function_VpcConfig `yaml:"VpcConfig,omitempty"`
	TracingConfig *properties.Function_TracingConfig `yaml:"TracingConfig,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
	Environment *properties.Function_Environment `yaml:"Environment,omitempty"`
	DeadLetterConfig *properties.Function_DeadLetterConfig `yaml:"DeadLetterConfig,omitempty"`
	Code *properties.Function_Code `yaml:"Code"`
}

func NewLambdaFunction(properties LambdaFunctionProperties, deps ...interface{}) LambdaFunction {
	return LambdaFunction{
		Type:       "AWS::Lambda::Function",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseLambdaFunction(name string, data string) (cf types.ValueMap, err error) {
	var resource LambdaFunction
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: LambdaFunction - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource LambdaFunction) Validate() []error {
	return resource.Properties.Validate()
}

func (resource LambdaFunctionProperties) Validate() []error {
	errs := []error{}
	if resource.Handler == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Handler'"))
	}
	if resource.Role == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Role'"))
	}
	if resource.Runtime == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Runtime'"))
	}
	if resource.Code == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'Code'"))
	} else {
		errs = append(errs, resource.Code.Validate()...)
	}
	return errs
}
