Ingress: HTTPS -> Nginx public ingress controller -> HTTP -> K8s services (Add App Gateway in future if needed)

Egress: K8s NetworkPolicies OR UDRs at subnet level + Azure Firewall on a Hub subnet

Ingress:
    - HTTPS Request ingress (wildcard) TLS cert created by cert-manager if doesnt exist, and traffic forwarded
    - Ingress is TLS termination point and forwards traffic to workload pods via HTTP
