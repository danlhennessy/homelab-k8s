package main

import (
	"context"
	"fmt"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/keyvault/azsecrets"
)

func Test() {
	fmt.Println("This is a secret function.")
}

func Get_Secret(mySecretName string) (string, error) {
	vaultURI := os.Getenv("AZURE_KEY_VAULT_URI")

	// Create a credential using the NewDefaultAzureCredential type.
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return "", fmt.Errorf("failed to obtain a credential: %v", err)
	}

	// Establish a connection to the Key Vault client
	client, err := azsecrets.NewClient(vaultURI, cred, nil)
	if err != nil {
		return "", fmt.Errorf("failed to create Key Vault client: %v", err)
	}

	// Get a secret. An empty string version gets the latest version of the secret.
	version := ""
	resp, err := client.GetSecret(context.Background(), mySecretName, version, nil)
	if err != nil {
		return "", fmt.Errorf("failed to get the secret: %v", err)
	}

	return *resp.Value, nil
}

func Get_Secrets(mySecretNames []string) (map[string]string, error) {
	vaultURI := os.Getenv("AZURE_KEY_VAULT_URI")

	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		return nil, fmt.Errorf("failed to obtain a credential: %v", err)
	}

	client, err := azsecrets.NewClient(vaultURI, cred, nil)
	if err != nil {
		return nil, fmt.Errorf("failed to create Key Vault client: %v", err)
	}

	// Initialize a map to store the secret values
	secretValues := make(map[string]string)

	// Iterate through the provided secret names
	for _, secretName := range mySecretNames {
		// Get a secret. An empty string version gets the latest version of the secret.
		version := ""
		resp, err := client.GetSecret(context.Background(), secretName, version, nil)
		if err != nil {
			return nil, fmt.Errorf("failed to get the secret '%s': %v", secretName, err)
		}

		// Add the secret value to the map
		secretValues[secretName] = *resp.Value
	}

	return secretValues, nil
}
