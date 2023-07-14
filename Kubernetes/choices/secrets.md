
# Secrets
Kubernetes secrets  (Move to cloud managed secrets solution when decided)
Usage:
    As files in a volume mounted on one or more of its containers.
    As container environment variable.
    By the kubelet when pulling images for the Pod.

Add secret keys as env variables in container:

apiVersion: v1
  kind: Pod
  metadata:
    name: envfrom-secret
  spec:
    containers:
    - name: envars-test-container
      image: nginx
      envFrom:
      - secretRef:
          name: test-secret

# Environment variables
Config Maps:
    data field for utf-8 strings , binaryData for base64-encoded strings
