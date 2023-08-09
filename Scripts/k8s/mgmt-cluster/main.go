package main

import (
	"fmt"
	"time"
)

func main() {

	Create_Namespaces([]string{
		"argocd",
		"crossplane-system",
	})

	// Grab secrets from AZ Key Vault
	azCreds, _ := Get_Secret("azure-credentials")
	sv_principal, _ := Get_Secrets([]string{"aks-clientSecret", "aks-clientID"})

	// Push secrets to Kubernetes (Current kubeconfig context)
	Apply_String_Secret("crossplane-system", "azure-secret", "creds", azCreds)
	Apply_Multi_String_Secret("default", "sv-principal", sv_principal)
	// Apply_JSON_Secret("default", "json-test", azCreds)

	Helm_Install("crossplane", "crossplane-stable/crossplane", "crossplane-system")

	Apply_Manifests("argocd", []string{
		"/media/dan/Transcend9/Backup/Work/DevOps/homelab/ArgoCD/argocd-stable.yaml",
		"/media/dan/Transcend9/Backup/Work/DevOps/homelab/ArgoCD/applicationsets/cluster-git-matrix.yaml",
	})
	// CRDs take ~ 40s to install
	fmt.Println("Waiting for CRDs.....")
	time.Sleep(60 * time.Second)
	fmt.Println("Continuing.....")

	Apply_Manifests("default", []string{
		"/media/dan/Transcend9/Backup/Work/DevOps/homelab/Kubernetes/crossplane/*.yaml",
	})
	// CRDs take ~ 70s to install
	fmt.Println("Waiting for CRDs.....")
	time.Sleep(90 * time.Second)
	fmt.Println("Continuing.....")

	Apply_Manifests("crossplane-system", []string{
		"/media/dan/Transcend9/Backup/Work/DevOps/homelab/Kubernetes/crossplane/az-provider-config/az-providerconfig.yaml",
		"/media/dan/Transcend9/Backup/Work/DevOps/homelab/Kubernetes/crossplane/cluster-infra-kubenet/*.yaml",
		"/media/dan/Transcend9/Backup/Work/DevOps/homelab/Kubernetes/crossplane/cluster-infra-kubenet/uk-west/*.yaml",
	})
}
