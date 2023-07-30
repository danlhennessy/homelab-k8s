# Steps to deploy management and workload clusters

## 1 - Management cluster bootstrap
az login
cd Scripts/k8s/mgmt-cluster && export AZURE_KEY_VAULT_URI=(Key Vault URI)
go run .

## 2 - Add additional clusters
Py script
Steps:
Add contexts to kubeconfig from each additional cluster (az aks get-credentials..)
Change context back to mgmt cluster:
    argocd login
    argocd cluster add ..
    sync applications in application set (Or set to autosync)