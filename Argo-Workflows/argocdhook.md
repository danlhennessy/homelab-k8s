Trigger a workflow using an ArgoCD Hook:

apiVersion: batch/v1
kind: Job
metadata:
  generateName: schema-migrate-
  annotations:
    argocd.argoproj.io/hook: PreSync  # <-- Add this notation to a manifest in the repo that argocd uses as source

# Hook Types

Hook 	    Description
PreSync 	Executes prior to the application of the manifests.
Sync 	    Executes after all PreSync hooks completed and were successful, at the same time as the application of the manifests.
Skip 	    Indicates to Argo CD to skip the application of the manifest.
PostSync 	Executes after all Sync hooks completed and were successful, a successful application, and all resources in a Healthy state.
SyncFail 	Executes when the sync operation fails. Available starting in v1.2

Basic Example:

apiVersion: argoproj.io/v1alpha1
kind: Workflow
metadata:
  generateName: kubectl-get-pods-
  namespace: argo
  annotations:
    argocd.argoproj.io/hook: Sync
spec:
  serviceAccountName: argo-admin
  entrypoint: kubectl-get-pods
  templates:
    - name: kubectl-get-pods
      container:
        image: bitnami/kubectl:latest
        command: [kubectl]
        args: [get, pods, -n, default]