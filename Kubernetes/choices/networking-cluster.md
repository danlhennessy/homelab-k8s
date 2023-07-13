AKS + kubenet + Istio + Network policies for pod to pod rules

or

AKS + Azure CNI w/ Cilium (Assign IP addresses from Virtual Network not Overlay Network)
(Currently no Hubble or gateway API - use regular services until these are natively supported by AKS)

or 

https://learn.microsoft.com/en-gb/azure/aks/use-byo-cni?tabs=azure-cli + manually install Cilium (High overheads)

or

AKS + kubenet + nginx/traefik