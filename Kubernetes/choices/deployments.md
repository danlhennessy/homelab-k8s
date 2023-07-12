# Deployment Stages

1. Application repos with their own app CI/CD , ending in a new image being pushed to a container registry

   - Sandbox docker + kind dev environment for quick tests

2. dev app manifests in repo dev branch. Prod/Staging app manifests in repo master branch. Test on dev branch , until ready to PR and merge dev into master. CI for manifests: Datree scan

3. ArgoCD monitors repos and when app manifests are updated, it will sync the relevant namespace/cluster: Dev, Staging and Prod Namespaces
   Argo Workflows triggered as a result of ArgoCD syncs - PreSync / Sync / PostSync / SyncFailed Hooks. Workflows runs specific and general checks
   PreSync Kyverno, PostSync Trivy / Datree / Falco
4. After merge to master - staging will build automatically, run tests/scans and send alerts using argo workflows.

5. Optional deploy to prod at this point - argo rollouts for canary deployment
   - AKS + Azure CNI with Cilium dataplane, Regular Ingress / Load Balancers + Argo Rollouts
   - When released for AKS: Argo Rollouts + Cilium API Gateway https://github.com/argoproj-labs/rollouts-plugin-trafficrouter-gatewayapi/

6. Monitor cluster/apps with post deploy tools. Manual rollback optional - argocd app rollback

# Dev Environment:

Pre image update:
- docker compose / docker run for pure local

Dev test:
Build docker image locally using docker compose --build or docker build.
Push image to container registry
Update development branch with that image tag (Multiple dev branches per team if needed)

ArgoCD syncs to Dev Cluster in their namespace.
How to deal with multiple tests at once  - service routing?

# Post Deploy Monitoring/Alerting/Security

Pixie: Network flow - DNS requests/TCP drops + retransits, HTTP latency, DB query latency, Node/Pod resource usage
Falcosidekick: Runtime security events -> Slack / ECK
Trivy: Vulnerability / Config / Compliance reports
APM: Pixie to start, add elastic APM for more detail (Requires app instrumentation)
Load Testing: K6