# kubeconfig
kubectl uses kubeconfig to locate and authenticate with a cluster API server
Authentication can be CA certificate based or username/password based.
Certificates are located in /etc/kubernetes/pki on the master node by default. On minikube: ~/.minikube/.
    Kubernetes will only accept certificates signed by kubernetes CA
    Managed kubernetes will use OATH instead (Certificate data + key and a refresh token)
When running kubectl outside a pod, it runs against the namespace define in kubeconfig - Default location = $HOME/.kube
KUBECONFIG env variable can hold a list of kubeconfig files, kubectl will merge these if it exists.
    Context: Group parameters together (Cluster, Namespace and User)

# ServiceAccount
Apps or Services will use service accounts to authenticate with the API server. Most apps will not need this, but anything that needs to integrate with the cluster will, like a CI/CD tool that sends a GET request to the API server or job that runs kubectl commands.
When running kubectl from within a pod, kubectl runs against the namespace of the pods serviceaccount.

# API Server
Azure managed endpoint, access controlled via Azure AD for users. For workloads uses serviceaccounts + managed identity/service principal to authenticate

## User
    Azure AD Authentication + Kubernetes RBAC - Azure AD Groups linked to Kubernetes Roles

    # Authentication
    az aks get-credentials --resource-group RG --name NAME
    Optional: kubelogin convert-kubeconfig -l azurecli

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
    WI uses kubelet to provide a token to a workload at a file path. The workload uses the token to authenticate via AzureAD with an Azure resource. https://azure.github.io/azure-workload-identity/docs/quick-start.html

    Linking a serviceaccount with a user assigned managed identity:

    apiVersion: v1
    kind: ServiceAccount
    metadata:
        annotations:
            azure.workload.identity/client-id: ${APPLICATION_CLIENT_ID:-$USER_ASSIGNED_IDENTITY_CLIENT_ID}
        name: ${SERVICE_ACCOUNT_NAME}
        namespace: ${SERVICE_ACCOUNT_NAMESPACE}
    
    Establish federated identity credential between the identity and the service account issuer & subject:
    
    az identity federated-credential create \
        --name "kubernetes-federated-credential" \
        --identity-name "${USER_ASSIGNED_IDENTITY_NAME}" \
        --resource-group "${RESOURCE_GROUP}" \
        --issuer "${SERVICE_ACCOUNT_ISSUER}" \
        --subject "system:serviceaccount:${SERVICE_ACCOUNT_NAMESPACE}:${SERVICE_ACCOUNT_NAME}"
