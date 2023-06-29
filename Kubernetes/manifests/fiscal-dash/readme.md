kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
kubectl apply -f monitor.yaml

Or `argocd app sync fiscal-dash`

Access app externally at nodeIP:nodeport , or internally at clusterIP:port

kubectl delete -f deployment.yaml
kubectl delete -f service.yaml
kubectl delete -f monitor.yaml