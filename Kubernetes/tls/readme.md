https://cert-manager.io/docs/tutorials/acme/nginx-ingress/

Create issuers in same namespace as the certs you need to create

kubectl apply -f https://github.com/cert-manager/cert-manager/releases/download/v1.13.1/cert-manager.yaml


helm repo add jetstack https://charts.jetstack.io
helm repo update

helm install \
cert-manager jetstack/cert-manager \
--namespace cert-manager \
--create-namespace \
--version v1.13.1 \
--set installCRDs=true

helm uninstall cert-manager -n cert-manager

## Add new ingress with TLS + automatic cert

1. Install Cert manager and create cluster issuer
1a. Allow port 80 inbound to cluster
2. Create Ingress resource with annotations: *cert-manager.io/issuer: "letsencrypt-staging"*. Or create a specific certificate resource
3. Cert Manager will create Cert, Order and Challenge resources to build a certificate secret (Describe each resource to check progress or view cert manager pod logs)
4. Troubleshooting: https://cert-manager.io/v1.6-docs/faq/acme/
5. New Cert should be created, signed and used by Ingress. To prevent ingress-nginx defaulting to a fake cert, adjust in manifest: - --default-ssl-certificate=default/example-cert

## Renew

Add a new dnsName to the cert resource and sync with argocd. A new challenge should be created to validate the dns route so it must lead to the cluster, then entire cert will be renewed
Check when cert last updated: kubectl get secret example-cert --show-managed-fields -o jsonpath='{range .metadata.managedFields[*]}{.manager}{" did "}{.operation}{" at "}{.time}{"\n"}{end}'