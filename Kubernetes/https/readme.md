# cert-manager + nginx ingress
https://cert-manager.io/docs/tutorials/acme/nginx-ingress/

helm install \
  cert-manager jetstack/cert-manager \
  --namespace cert-manager \
  --create-namespace \
  --version v1.12.0 \
  --set installCRDs=true

- Use DNS to send wildcard DNS name to ingress controller IP
- Create Issuers - Kubernetes/https/cert-manager
- Create Ingress resource with annotation: `cert-manager.io/issuer: "letsencrypt-staging"` for testing, then `cert-manager.io/issuer: "letsencrypt-prod"` when working
- Cert resource will be autocreated and requested using letsencrypt and your chosen DNS name in the ingress
- If requested cert is successful, a secret will be created with the name chosen in the ingress