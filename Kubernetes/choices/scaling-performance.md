Cluster:
    Node scaling with AKS autoscaler (Horizontal - automatically creates additional nodes based on pod demand. Can configure triggers in cluster autoscaler profile)

Pod:
    KEDA - scalers: Prometheus, PostgreSQL etc
    Horizontal: HorizontalPodAutoscaler https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/
    Vertical: https://learn.microsoft.com/en-us/azure/aks/vertical-pod-autoscaler
    Only use one in a single node pool, prefer horizontal where suitable.

# Compute resource - https://kubernetes.io/docs/concepts/configuration/manage-resources-containers/
Containers are burstable by default and will use up to the available compute/memory capacity of the node unless limits/requests are defined in their spec.
Prefer only requests to limits unless required. https://home.robusta.dev/blog/stop-using-cpu-limits
KRR can help determine appropriate resource numbers: https://github.com/robusta-dev/krr