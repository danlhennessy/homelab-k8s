Chaos Engineering - resiliency testing

# Resilience

Failover at the service level, not at ingress. Serviceroutes to pods based on health checks. Can add additional custom liveness/readiness probes to apps to take advantage of this.
- https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes

Use liveness probes in cases where restarting may help, readiness probes if app receives network traffic but may not be immediately ready to accept.
https://developers.redhat.com/blog/2020/11/10/you-probably-need-liveness-and-readiness-probes#example_1__a_static_file_server__nginx_ for more detail

Example livenessprobes:
    livenessProbe:
      exec:
        command:
        - cat
        - /tmp/healthy
      initialDelaySeconds: 5
      periodSeconds: 5
    livenessProbe:
      httpGet:
        path: /healthz
        port: 8080
        httpHeaders:
        - name: Custom-Header
          value: Awesome
      initialDelaySeconds: 3
      periodSeconds: 3

Pod anti-affinity - spread newly scheduled pods across nodes to improve resilience. Example of required and preferred affinities:

apiVersion: v1
kind: Pod
metadata:
  name: with-node-affinity
spec:
  affinity:
    nodeAffinity:
      requiredDuringSchedulingIgnoredDuringExecution:
        nodeSelectorTerms:
        - matchExpressions:
          - key: topology.kubernetes.io/zone
            operator: In
            values:
            - antarctica-east1
            - antarctica-west1
      preferredDuringSchedulingIgnoredDuringExecution:
      - weight: 1
        preference:
          matchExpressions:
          - key: another-node-label-key
            operator: In
            values:
            - another-node-label-value

# Disaster Recovery

GitOps + Velero for stateful workloads