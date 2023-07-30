package main

func main() {
	// Grab secrets from AZ Key Vault
	azCreds, _ := Get_Secret("azure-credentials")
	sv_principal, _ := Get_Secrets([]string{"aks-clientSecret", "aks-clientID"})

	// Push secrets to Kubernetes (Current kubeconfig context)
	Apply_String_Secret("crossplane-system", "azure-secret", "creds", azCreds)
	Apply_Multi_String_Secret("default", "sv-principal", sv_principal)
	// Apply_JSON_Secret("default", "json-test", azCreds)

	manifests := []string{
		"/media/dan/Transcend9/Backup/Work/DevOps/homelab/Kubernetes/manifests/apache/apache.yaml",
		"/media/dan/Transcend9/Backup/Work/DevOps/homelab/Kubernetes/manifests/apache/service.yaml",
	}
	Apply_Manifests(manifests)
}
