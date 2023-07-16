AKS + kubenet + Istio + Network policies for pod to pod rules

or

AKS + Azure CNI w/ Cilium (Assign IP addresses from Virtual Network not Overlay Network)
(Currently no Hubble or gateway API - use regular services until these are natively supported by AKS)

or 

https://learn.microsoft.com/en-gb/azure/aks/use-byo-cni?tabs=azure-cli + manually install Cilium (High overheads)

or

AKS + kubenet + nginx/traefik

# API Traffic
Public/Private cluster w/ API Integration https://techcommunity.microsoft.com/t5/core-infrastructure-and-security/public-and-private-aks-clusters-demystified/ba-p/3716838

# Egress
NetworkPolicies