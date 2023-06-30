# Deployment Stages

1. Application repos with their own app CI/CD , ending in a new image being pushed to a container registry

   - Sandbox docker setup for quick tests

2. dev app manifests in repo dev branch. Prod/Staging app manifests in repo master branch. Test on dev branch , until ready to PR and merge dev into master.

3. ArgoCD monitors repos and when app manifests are updated, it will sync the relevant namespace/cluster: Dev, Staging and Prod Namespaces
   Argo Workflows triggered as a result of ArgoCD syncs - PreSync / Sync / PostSync / SyncFailed Hooks. Workflows runs specific and general checks

4. After merge to master - staging will build automatically, run tests/scans and send alerts using argo workflows.

5. Optional deploy to prod at this point - argo rollouts for blue/green deployment / canary deployment?


# Choices

