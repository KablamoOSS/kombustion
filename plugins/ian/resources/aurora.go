// +build plugin

package resources

import (
	"log"
	"strconv"

	"github.com/KablamoOSS/kombustion/pluginParsers/resources"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

type AuroraConfig struct {
	Properties struct {
		InstanceCount *int    `yaml:"InstanceCount" example:"1"`
		Username      *string `yaml:"Username" example:"admin"`
		Password      *string `yaml:"Password" example:"Password123!"`
		InstanceType  *string `yaml:"InstanceType"`
	} `yaml:"Properties"`
}

func ParseAurora(name string, data string) (cf types.ValueMap, err error) {
	// Parse the config data
	var config AuroraConfig
	if err = yaml.Unmarshal([]byte(data), &config); err != nil {
		return
	}

	// validate the config
	config.Validate()

	// create a group of objects (each to be validated)
	cf = make(types.ValueMap)

	// defaults
	instanceCount := 1
	if config.Properties.InstanceCount != nil {
		instanceCount = *config.Properties.InstanceCount
	}

	dbInstanceClass := "db.t2.small"
	if config.Properties.InstanceType != nil {
		dbInstanceClass = *config.Properties.InstanceType
	}

	// create the cluster
	cf[name+"AuroraCluster"] = resources.NewRDSDBCluster(
		resources.RDSDBClusterProperties{
			Engine:             "aurora",
			MasterUsername:     *config.Properties.Username,
			MasterUserPassword: *config.Properties.Password,
		},
	)

	// create the cluster members (instances)
	for i := 1; i <= instanceCount; i++ {
		cf[name+"AuroraInstance"+strconv.Itoa(i)] = resources.NewRDSDBInstance(
			resources.RDSDBInstanceProperties{
				DBClusterIdentifier: map[string]string{"Ref": (name + "AuroraCluster")},
				DBInstanceClass:     dbInstanceClass,
				Engine:              "aurora",
			},
		)
	}

	return
}

// Validate - input Config validation
func (this AuroraConfig) Validate() {
	if this.Properties.Username == nil || this.Properties.Password == nil {
		log.Println("WARNING: AuroraConfig - Missing 'Username' and/or 'Password'")
	}
}
