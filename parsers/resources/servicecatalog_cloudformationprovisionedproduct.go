package resources

import (
	yaml "github.com/KablamoOSS/yaml"
	"github.com/KablamoOSS/kombustion/types"
	"log"
)

type ServiceCatalogCloudFormationProvisionedProduct struct {
	Type       string                      `yaml:"Type"`
	Properties ServiceCatalogCloudFormationProvisionedProductProperties `yaml:"Properties"`
	Condition  interface{}                 `yaml:"Condition,omitempty"`
	Metadata   interface{}                 `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                 `yaml:"DependsOn,omitempty"`
}

type ServiceCatalogCloudFormationProvisionedProductProperties struct {
	AcceptLanguage interface{} `yaml:"AcceptLanguage,omitempty"`
	PathId interface{} `yaml:"PathId,omitempty"`
	ProductId interface{} `yaml:"ProductId,omitempty"`
	ProductName interface{} `yaml:"ProductName,omitempty"`
	ProvisionedProductName interface{} `yaml:"ProvisionedProductName,omitempty"`
	ProvisioningArtifactId interface{} `yaml:"ProvisioningArtifactId,omitempty"`
	ProvisioningArtifactName interface{} `yaml:"ProvisioningArtifactName,omitempty"`
	NotificationArns interface{} `yaml:"NotificationArns,omitempty"`
	ProvisioningParameters interface{} `yaml:"ProvisioningParameters,omitempty"`
	Tags interface{} `yaml:"Tags,omitempty"`
}

func NewServiceCatalogCloudFormationProvisionedProduct(properties ServiceCatalogCloudFormationProvisionedProductProperties, deps ...interface{}) ServiceCatalogCloudFormationProvisionedProduct {
	return ServiceCatalogCloudFormationProvisionedProduct{
		Type:       "AWS::ServiceCatalog::CloudFormationProvisionedProduct",
		Properties: properties,
		DependsOn:  deps,
	}
}

func ParseServiceCatalogCloudFormationProvisionedProduct(name string, data string) (cf types.ValueMap, err error) {
	var resource ServiceCatalogCloudFormationProvisionedProduct
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ServiceCatalogCloudFormationProvisionedProduct - ", err)
		}
		return
	}
	cf = types.ValueMap{name: resource}
	return
}

func (resource ServiceCatalogCloudFormationProvisionedProduct) Validate() []error {
	return resource.Properties.Validate()
}

func (resource ServiceCatalogCloudFormationProvisionedProductProperties) Validate() []error {
	errs := []error{}
	return errs
}
