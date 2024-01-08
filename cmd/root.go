package commands

import (
	"andy/constants"

	"github.com/spf13/cobra"
)

// type Meta struct {
// 	Version string `yaml:"version"`
// }

// func getVersion() string {
// 	var meta Meta

// 	metaFile, _ := filepath.Abs("./configs/meta.yaml")
// 	yamlFile, yamlFileErr := os.ReadFile(metaFile)

// 	if yamlFileErr != nil {
// 		panic(yamlFileErr)
// 	}

// 	if err := yaml.Unmarshal(yamlFile, &meta); err != nil {
// 		panic(err)
// 	}

// 	return meta.Version
// }

var RootCmd = &cobra.Command{
	Short: "Hi! I'm Andy" + constants.Logo + "!",
	CompletionOptions: cobra.CompletionOptions{
		DisableDefaultCmd: true,
	},
	// Version: getVersion(),
	Version: "0.0.2",
}
