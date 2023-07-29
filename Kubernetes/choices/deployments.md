# Deployment Stages

1. Application repos with their own app CI/CD , ending in a new image being pushed to a container registry

   - Sandbox docker + kind/minikube/rancher desktop dev environment for quick tests

2. dev app manifests in repo dev branch. Prod/Staging app manifests in repo master branch. Test on dev branch , until ready to PR and merge dev into master. CI for manifests: Datree scan

3. ArgoCD monitors repos and when app manifests are updated, it will sync the relevant namespace/cluster: Dev, Staging and Prod Namespaces. Manifest ordered release: bootstrap -> CRDs -> config/secret setup -> applications
   Argo Workflows triggered as a result of ArgoCD syncs - PreSync / Sync / PostSync / SyncFailed Hooks. Workflows runs specific and general checks
   PreSync Kyverno, PostSync Trivy / Datree / Falco
4. After merge to master - staging will build automatically, run tests/scans and send alerts using argo workflows.

5. Optional deploy to prod at this point - argo rollouts for canary deployment
   - AKS + Azure CNI with Cilium dataplane, Regular Ingress / Load Balancers + Argo Rollouts
   - When released for AKS: Argo Rollouts + Cilium API Gateway https://github.com/argoproj-labs/rollouts-plugin-trafficrouter-gatewayapi/

6. Monitor cluster/apps with post deploy tools. Manual rollback optional - argocd app rollback

# Dev Environment:

Pre image update:
- docker compose / docker run for pure local. Docker compose -> kompose -> k8s manifests
- DevSpace synced with local or AKS cluster (1 devspace session per namespace). devspace dev deploys pod to namespace and devspace purge tears down

ArgoCD syncs to Dev Cluster in their namespace.
How to deal with multiple tests at once  - multiple namespaces + devspace.yaml files:

`Running devspace dev multiple times in parallel for the same project does not work and returns a fatal error because multiple instances of port forwarding and file sync will disturb each other. However, running devspace dev for several different projects (i.e. 2 different devspace.yaml files with different name) works seamlessly.`

# Dev test:
Build docker image locally using docker compose --build or docker build.
Push image to container registry
Update development branch with that image tag (Multiple dev branches per team if needed)
# Post Deploy Monitoring/Alerting/Security

Pixie: Network flow - DNS requests/TCP drops + retransits, HTTP latency, DB query latency, Node/Pod resource usage
Falcosidekick: Runtime security events -> Slack / ECK
Trivy: Vulnerability / Config / Compliance reports
APM: Pixie to start, add elastic APM for more detail (Requires app instrumentation)
Load Testing: K6