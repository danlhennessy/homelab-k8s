Cluster:
    Node scaling with AKS autoscaler (Horizontal - automatically creates additional nodes based on pod demand. Can configure triggers in cluster autoscaler profile)

Pod:
    Horizontal: HorizontalPodAutoscaler https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale/
    Vertical: https://learn.microsoft.com/en-us/azure/aks/vertical-pod-autoscaler
    Only use one in a single node pool, prefer horizontal where suitable.