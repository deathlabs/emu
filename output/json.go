package output

import (
	"encoding/json"
	"fmt"
	"os"
)

func ToJSON(data interface{}) {
	var (
		err      error
		jsonData []byte
	)

	jsonData, err = json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Printf("Error formatting JSON: %v\n", err)
		os.Exit(1)
	}

	fmt.Println(string(jsonData))
}
