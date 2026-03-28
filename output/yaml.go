package output

import (
	"fmt"
	"os"

	"go.yaml.in/yaml/v3"
)

func ToYAML(data interface{}) {
	var (
		err      error
		yamlData []byte
	)

	yamlData, err = yaml.Marshal(data)
	if err != nil {
		fmt.Printf("Error formatting YAML: %v\n", err)
		os.Exit(1)
	}

	fmt.Print(string(yamlData))
}
