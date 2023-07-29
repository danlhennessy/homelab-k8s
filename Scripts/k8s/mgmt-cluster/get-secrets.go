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
