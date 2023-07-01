# Istio
istioctl install --set profile=demo -y

kubectl label namespace default istio-injection=enabled
kubectl label namespace fiscal-ns istio-injection=enabled
kubectl label namespace monitoring istio-injection=enabled

- Deploy a pod with a service to default namespace and an Istio sidecar pod will be deployed along with it

Teardown: 
kubectl delete -f samples/addons
istioctl uninstall -y --purge
kubectl delete namespace istio-system
kubectl label namespace default istio-injection-

# Kiali
<!-- helm install \
  --namespace istio-system \
  --set auth.strategy="anonymous" \
  --repo https://kiali.org/helm-charts \
  -f kiali-values.yaml \
  kiali-server \
  kiali-server -->

helm install \
  --namespace istio-system \
  --set auth.strategy="anonymous" \
  --repo https://kiali.org/helm-charts \
  kiali-server \
  kiali-server

kubectl apply -f https://raw.githubusercontent.com/istio/istio/release-1.18/samples/addons/prometheus.yaml

istioctl dashboard kiali

Working Istio metrics + visualisations:
- Install Istio, Kiali and Prometheus addon as above
- Label namespace with istio-injection=enabled
- Deploy app and make sure Istio sidecar is running
- Check Kiali graphs for traffic / metrics
# Teardown

helm uninstall --namespace istio-system kiali-server
kubectl delete -f https://raw.githubusercontent.com/istio/istio/release-1.18/samples/addons/prometheus.yaml
