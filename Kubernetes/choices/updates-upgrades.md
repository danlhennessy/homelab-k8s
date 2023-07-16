AKS: https://learn.microsoft.com/en-us/azure/architecture/reference-architectures/containers/aks/baseline-aks#security-updates
    

Velero + Kured

Kubernetes Version Upgrade: https://learn.microsoft.com/en-us/azure/aks/upgrade-cluster?tabs=azure-cli
    K8s version upgrade to be tested on 1 cluster for 1 week, then added to other.

Node Image upgrade: https://learn.microsoft.com/en-us/azure/aks/node-image-upgrade
    Check weekly for node image upgrades with:
    az aks nodepool get-upgrades \
    --nodepool-name mynodepool \
    --cluster-name myAKSCluster \
    --resource-group myResourceGroup
    Node image upgrade to be tested on 1 cluster for 1 week, then added to other.

OS/Runtime Patches (With Reboots): Kured