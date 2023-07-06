helm repo add okteto https://charts.okteto.com
helm repo update
helm install okteto okteto/okteto -f config.yaml --namespace=okteto --create-namespace

sudo kubectl port-forward service/okteto-ingress-nginx-controller 443:443 --namespace okteto
https://okteto.localtest.me/login#token=******

okteto context use https://okteto.localtest.me --token ****** --insecure-skip-tls-verify

helm delete okteto -n okteto