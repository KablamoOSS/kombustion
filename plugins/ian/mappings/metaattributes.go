// +build plugin

package mappings

import (
	"io/ioutil"
	"net/http"

	"github.com/KablamoOSS/kombustion/types"
)

func ParseMetaAttributes(name string, data string) (cf types.ValueMap, err error) {
	// create a group of objects (each to be validated)
	cf = make(types.ValueMap)

	ip := "unknown"

	resp, err := http.Get("https://api.ipify.org/")
	if err == nil {
		body, err := ioutil.ReadAll(resp.Body)
		if err == nil {
			ip = string(body)
		}
	}
	defer resp.Body.Close()

	cf["MetaAttributes"] = map[string]map[string]string{
		"Networking": map[string]string{
			"PublicIp": ip,
		},
	}

	return
}
