Istio Ingress controller - gateways per namespace and virtual services -> services

/ ebpf - cilium?

Network policies for pod to pod rules

az aks create -n <clusterName> -g <resourceGroupName> -l <location> \
  --network-plugin azure \
  --network-plugin-mode overlay \
  --pod-cidr 192.168.X.X/16 \
  --network-dataplane cilium
(No Hubble currently)

or 

https://learn.microsoft.com/en-gb/azure/aks/use-byo-cni?tabs=azure-cli + manually install Cilium
