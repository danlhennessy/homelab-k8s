
istioctl install --set profile=demo -y
Working Istio metrics + visualisations:
- Install Istio, Kiali and Prometheus addon as below
- Label namespace with istio-injection=enabled
- Deploy app and make sure Istio sidecar is running
- Check Kiali graphs for traffic / metrics
# Label namespaces
kubectl create ns nginx-staging
kubectl label namespace nginx-staging istio-injection=enabled

# Helm
helm repo add istio https://istio-release.storage.googleapis.com/charts
helm repo update

helm install istio-base istio/base -n istio-system --create-namespace --set defaultRevision=default
helm install istiod istio/istiod -n istio-system -f values/sidecar-exempt.yaml

helm install istio-ingress istio/gateway -n istio-ingress --create-namespace

helm install --namespace istio-system --set auth.strategy="anonymous" --repo https://kiali.org/helm-charts kiali-server kiali-server
istioctl dashboard kiali

kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.18/samples/addons/prometheus.yaml
kubectl apply -f gateway.yaml
kubectl apply -f virtualsvc.yaml

# Helm Teardown

helm delete istio-ingress -n istio-ingress
kubectl delete namespace istio-ingress
helm delete istiod -n istio-system
helm delete istio-base -n istio-system
kubectl delete namespace istio-system

helm uninstall --namespace istio-system kiali-server
kubectl delete -f https://raw.githubusercontent.com/istio/istio/release-1.18/samples/addons/prometheus.yaml
