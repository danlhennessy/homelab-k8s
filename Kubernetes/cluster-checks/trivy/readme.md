https://blog.aquasec.com/kubernetes-cluster-security-with-trivy
https://aquasecurity.github.io/trivy/dev/tutorials/kubernetes/cluster-scanning/

Trivy Operator to automatically scan resources every 6 hours

Installed via ArgoCD - trivy-operator.yaml

# Reports

kubectl get vulnerabilityreports --all-namespaces -o wide
kubectl get configauditreports --all-namespaces -o wide