# Mesh addon
az aks mesh enable --resource-group 101-sandbox-343lb0d1 --name sandboxcluster
az aks show --resource-group 101-sandbox-343lb0d1 --name sandboxcluster  --query 'serviceMeshProfile.mode'

az aks mesh disable --resource-group 101-sandbox-343lb0d1 --name sandboxcluster

# AKS Labelling
kubectl label namespace default istio.io/rev=asm-1-17

# Ingress Gateway

az aks mesh enable-ingress-gateway --resource-group 101-sandbox-343lb0d1 --name sandboxcluster --ingress-gateway-type external

az aks mesh enable-ingress-gateway --resource-group 101-sandbox-343lb0d1 --name sandboxcluster --ingress-gateway-type internal
az aks show --resource-group --resource-group 101-sandbox-343lb0d1 --name sandboxcluster --query provisioningState


az aks mesh disable-ingress-gateway --resource-group 101-sandbox-343lb0d1 --name sandboxcluster  --ingress-gateway-type external


kubectl apply -f /media/dan/Transcend5/Backup/Work/DevOps/homelab/Kubernetes/manifests/nginx-basic/dev/deployment.yml
kubectl apply -f /media/dan/Transcend5/Backup/Work/DevOps/homelab/Kubernetes/manifests/nginx-basic/dev/service.yml

kubectl apply -f /media/dan/Transcend5/Backup/Work/DevOps/homelab/Kubernetes/service-mesh/ingress/gateway.yaml
kubectl apply -f /media/dan/Transcend5/Backup/Work/DevOps/homelab/Kubernetes/service-mesh/ingress/virtualsvc.yaml

# TLS Ingress Gateway - put secret, gateway and virtual svc in same ns as the ingress gateway
kubectl create secret tls ingress-cert --cert=fullchain.pem --key=privkey.pem -n aks-istio-ingress  (cert and key for host linked to ingress controller)
kubectl apply -f /media/dan/Transcend5/Backup/Work/DevOps/homelab/Kubernetes/service-mesh/ingress/httpsgateway.yaml -n aks-istio-ingress
kubectl apply -f /media/dan/Transcend5/Backup/Work/DevOps/homelab/Kubernetes/service-mesh/ingress/https-vs.yaml -n aks-istio-ingress

curl external IP

kubectl delete -f /media/dan/Transcend5/Backup/Work/DevOps/homelab/Kubernetes/manifests/nginx-basic/dev/deployment.yml
kubectl delete -f /media/dan/Transcend5/Backup/Work/DevOps/homelab/Kubernetes/manifests/nginx-basic/dev/service.yml