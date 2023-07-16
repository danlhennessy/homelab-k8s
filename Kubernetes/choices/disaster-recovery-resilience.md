Chaos Engineering - resiliency testing

# Resilience

Failover at the service level, not at ingress. Serviceroutes to pods based on health checks. Can add additional custom liveness/readiness probes to apps to take advantage of this.
- https://kubernetes.io/docs/tasks/configure-pod-container/configure-liveness-readiness-startup-probes/
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

# Disaster Recovery

GitOps + Velero for stateful workloads