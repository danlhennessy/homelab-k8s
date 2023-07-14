# Authentication
az aks get-credentials --resource-group RG --name NAME
kubelogin convert-kubeconfig -l azurecli

# Authorisation
Azure AD users + groups -> Kubernetes roles + rolebindings (Per namespace)
Role bindings like a kubernetes role e.g. admin to an Azure AD group
- Role : Namespace specific permissions
- ClusterRole : Cluster-wide permissions
- RoleBinding : Bind a role/clusterrole to a specific namespace
- ClusterRoleBinding : Bind a role/clusterrole to entire cluster


# Service accounts
Only for services / apps