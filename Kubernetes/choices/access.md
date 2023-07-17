## User
    Azure AD Authentication + Kubernetes RBAC - Azure AD Groups linked to Kubernetes Roles

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

## Workloads - https://learn.microsoft.com/en-us/azure/aks/use-managed-identity#summary-of-managed-identities

    A managed identity is a feature that provides an identity for an AKS cluster to interact with other Azure services securely. Alternative is service principal which needs to be manually updated.
    AKS default managed identities are for the control plane and the kubelet.
    e.g. Add acrpull role to kubelet managed identity to allow it to pull from Azure container registry

    Workloads inside AKS also need a way of authenticating themselves with Azure resources. Apps can use libraries e.g. MSAL or Azure workload identity - https://azure.github.io/azure-workload-identity/docs/

    Sometimes will need to authenticate as the managed identity before having the associated access of the MI.