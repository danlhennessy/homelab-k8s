External:
    Just one wildcard DNS record to Application gateway (e.g. *.gate.example.com). Based on the prefix, Ingress will route the request to the relevant service

Internal: https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/
    Pods and Services are assigned DNS names by kube-dns(CoreDNS)
    They can be queried with following formats:
    Service FQDN:   my-svc.my-namespace.svc.cluster-domain.example.
    Pod FQDN:   pod-ip-address.my-namespace.pod.cluster-domain.example.

    Relative DNS names can be used e.g.:
        To query a service in the same namespace: curl my-svc:port
        To query a service in another namespace: curl my-svc.namespace:port

DNS cache: https://kubernetes.io/docs/tasks/administer-cluster/nodelocaldns/