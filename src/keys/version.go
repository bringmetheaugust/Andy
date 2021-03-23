package keys

import (
	"andy/src/utils"
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type version struct {
	Version string `yaml:"version"`
	Author  string `yaml:"author"`
}

// Version show package version
func Version() {
	var version version
	yamlFile, err := ioutil.ReadFile("./version.yaml")

	if err != nil {
		utils.PanicMessage()
	}

	yaml.Unmarshal([]byte(yamlFile), &version)

	fmt.Println("version🏷 :   ", version.Version)
	fmt.Println("author🥇:    ", version.Author)
}
