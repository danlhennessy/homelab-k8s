# Steps to deploy management and workload clusters

## 1 - Management cluster bootstrap
az login
cd Scripts/k8s/mgmt-cluster && export AZURE_KEY_VAULT_URI=(Key Vault URI)
go run .

## 2 - Configure ArgoCD
