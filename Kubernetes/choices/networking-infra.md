Azure Traffic Manager -> 2 Clusters -> Ingress controller public IP

Ingress: HTTPS -> Nginx public ingress controller -> HTTP -> K8s services (Add App Gateway in future if needed)

Egress: K8s NetworkPolicies OR UDRs at subnet level + Azure Firewall on a Hub subnet

Ingress:
    - NSG at traffic manager level
    - HTTPS Request ingress (wildcard) TLS cert created by cert-manager if doesnt exist, and traffic forwarded
    - Ingress is TLS termination point and forwards traffic to workload pods via HTTP

# Front Door
1 Origin Group
2 Origins: Public IPs of ingress controller in each AKS cluster
Health probes enabled

# Traffic Manager
Cheaper, Simpler
Endpoint type = external endpoint
In testing - all requests go to working endpoint whereas Front Door still routed some to broken endpoint