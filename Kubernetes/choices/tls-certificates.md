1. Individual PFX certs stored in App Gateway or referenced in App Gateway from secrets management tool (e.g. x.example.com , y.example.com)
2. Wildcard cert (*.ingress.example.com) stored in K8s cluster and referenced by Ingress resources (Can be autocreated by cert-manager)
   
e.g: 

apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: apache-ingress
  annotations:
    kubernetes.io/ingress.class: "nginx"
    nginx.ingress.kubernetes.io/rewrite-target: /
    cert-manager.io/issuer: "letsencrypt-prod"
spec:
  tls:
  - hosts:
    - apache.ingress.danlhennessy.com
    secretName: apache-cert
  rules:
    - host: apache.ingress.danlhennessy.com
      http:
        paths:
          - path: /
            pathType: Prefix
            backend:
              service:
                name: apache-service
                port:
                  number: 8083