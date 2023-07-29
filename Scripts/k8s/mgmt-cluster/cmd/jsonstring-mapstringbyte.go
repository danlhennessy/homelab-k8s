package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	value, _ := Get_Secret("azure-credentials")

	// Convert the JSON string to a map[string]string
	var data map[string]string
	err := json.Unmarshal([]byte(value), &data)
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
		return
	}
	// Convert the string values to byte arrays
	secretData := make(map[string][]byte)
	for key, value := range data {
		secretData[key] = []byte(value)
	}

	CreateKubernetesSecret("default", "az-creds", secretData)
}
