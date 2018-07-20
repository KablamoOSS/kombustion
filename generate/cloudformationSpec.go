package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

var cfnEndpoints = map[string]string{
	"Sydney":           "https://d2stg8d246z9di.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json",
	"Singapore":        "https://doigdx0kgq9el.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json",
	"Mumbai":           "https://d2senuesg1djtx.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json",
	"Seoul":            "https://d1ane3fvebulky.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json",
	"Tokyo":            "https://d33vqc0rt9ld30.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json",
	"Canada":           "https://d2s8ygphhesbe7.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json",
	"Frankfurt":        "https://d1mta8qj7i28i2.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json",
	"London":           "https://d1742qcu2c1ncx.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json",
	"Ireland":          "https://d3teyb21fexa9r.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json",
	"Sao Paulo":        "https://d3c9jyj3w509b0.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json",
	"North Virginia":   "https://d1uauaxba7bl26.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json",
	"Ohio":             "https://dnwj8swjjbsbt.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json",
	"North California": "https://d68hl49wbnanq.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json",
	"Oregon":           "https://d201a2mn26r7lk.cloudfront.net/latest/gzip/CloudFormationResourceSpecification.json",
}

// getCloudformationSpecification downloads the latest cloudformation specification
func getCloudformationSpecification(sourceDir string, specList map[string]string) {

	clearSourceDir()

	for _, region := range sortSpecList(specList) {
		url := specList[region]
		cfnData := fetchSourceData(region)
		if len(cfnData) > 0 {
			log.Println("Downloaded region: " + region + " from " + url)
		}
	}
}

// Remove existing CFN Spec files
func clearSourceDir() {
	os.RemoveAll(sourceDir)

	os.Mkdir(sourceDir, 0744)
}

// Create a de-duplicated struct of the entire CFN Spec
func buildUniqueSet(sourceDir string, specList map[string]string) CfnSpec {
	uniquecfnData := CfnSpec{
		PropertyTypes: map[string]CfnType{},
		ResourceTypes: map[string]CfnType{},
	}
	for _, region := range sortSpecList(specList) {
		var tempcfnSpec CfnSpec
		cfnData, err := ioutil.ReadFile(fmt.Sprintf("%v%v.json", sourceDir, region))
		if err == nil {
			err = json.Unmarshal(cfnData, &tempcfnSpec)
			if err == nil {
				for k, v := range tempcfnSpec.PropertyTypes {
					if _, ok := uniquecfnData.PropertyTypes[k]; ok {
					} else {
						uniquecfnData.PropertyTypes[k] = v
					}
				}
				for k, v := range tempcfnSpec.ResourceTypes {
					if _, ok := uniquecfnData.ResourceTypes[k]; ok {
					} else {
						uniquecfnData.ResourceTypes[k] = v
					}
				}
			}
		}
	}
	return uniquecfnData
}

func fetchSourceData(region string) []byte {
	cfnURL := cfnEndpoints[region]
	request, err := http.NewRequest("GET", cfnURL, nil)
	checkError(err)

	request.Header.Set("Content-Type", "application/json")
	response, err := new(http.Client).Do(request)
	checkError(err)

	defer response.Body.Close()
	if response.StatusCode != http.StatusOK {
		log.Fatalf("Failed to fetch cfn template data: %v", response.StatusCode)
	}

	cfnData, err := ioutil.ReadAll(response.Body)
	checkError(err)

	err = ioutil.WriteFile(fmt.Sprintf("%v%v.json", sourceDir, region), cfnData, 0644)
	checkError(err)

	return cfnData
}
