# Authentication
az aks get-credentials --resource-group RG --name NAME
kubelogin convert-kubeconfig -l azurecli

# Authorisation
Azure AD users + groups -> Kubernetes roles + rolebindings (Per namespace)


# Service accounts
Only for services / apps

# Secrets
Kubernetes secrets  (Move to cloud managed secrets solution when decided)

# Environment variables
Config Maps:
    data field for utf-8 strings , binaryData for base64-encoded strings