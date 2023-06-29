istioctl install --set profile=demo -y

kubectl label namespace default istio-injection=enabled

- Deploy a pod with a service to default namespace and an Istio sidecar pod will be deployed along with it

Teardown: 
kubectl delete -f samples/addons
istioctl uninstall -y --purge
kubectl delete namespace istio-system
kubectl label namespace default istio-injection-