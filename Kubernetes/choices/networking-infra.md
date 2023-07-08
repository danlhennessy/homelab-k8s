Ingress: HTTPS only -> Azure App Gateway Public IP (With TLS Cert) -> Istio private ingress gateway -> HTTPS only -> K8s services

Egress: K8s NetworkPolicies OR UDRs at subnet level + Azure Firewall on a Hub subnet