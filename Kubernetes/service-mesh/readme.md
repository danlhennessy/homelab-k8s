Working Istio metrics + visualisations:
- Install Istio, Kiali and Prometheus addon as below
- Label namespace with istio-injection=enabled
- Deploy app and make sure Istio sidecar is running
- Check Kiali graphs for traffic / metrics
# Label namespaces
kubectl create ns nginx-staging
kubectl label namespace nginx-staging istio-injection=enabled

# Istioctl - https://istio.io/v1.6/docs/setup/install/istioctl/

`istioctl install -f values/staticip.yaml` (static private IP 10.96.0.4)
helm install --namespace istio-system --set auth.strategy="anonymous" --repo https://kiali.org/helm-charts kiali-server kiali-server

kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.18/samples/addons/prometheus.yaml
kubectl apply -f ingress/gateway.yaml
kubectl apply -f ingress/virtualsvc.yaml

Teardown: `istioctl uninstall --purge -y`
helm uninstall --namespace istio-system kiali-server
kubectl delete -f https://raw.githubusercontent.com/istio/istio/release-1.18/samples/addons/prometheus.yaml

# Helm
helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update

helm install istio-base istio/base -n istio-system --create-namespace --set defaultRevision=default
helm install istiod istio/istiod -n istio-system -f values/sidecar-exempt.yaml

helm install istio-ingress istio/gateway -n istio-ingress --create-namespace -f istio-ingressgateway.yaml
helm install istio-ingress istio/gateway -n istio-system --set gateways.istio-ingressgateway.loadBalancerIP=1.4.3.5

helm install --namespace istio-system --set auth.strategy="anonymous" --repo https://kiali.org/helm-charts kiali-server kiali-server
istioctl dashboard kiali

kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.18/samples/addons/prometheus.yaml
kubectl apply -f ingress/gateway.yaml
kubectl apply -f ingress/virtualsvc.yaml

# Helm Teardown

helm delete istio-ingress -n istio-ingress
kubectl delete namespace istio-ingress
helm delete istiod -n istio-system
helm delete istio-base -n istio-system
kubectl delete namespace istio-system

helm uninstall --namespace istio-system kiali-server
kubectl delete -f https://raw.githubusercontent.com/istio/istio/release-1.18/samples/addons/prometheus.yaml