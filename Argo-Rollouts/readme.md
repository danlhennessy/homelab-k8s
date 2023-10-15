## ArgoCD + Argo Rollouts Process (Success)
1. Version N runs on the cluster as a Rollout (managed by Argo CD). The Git repository is updated with version N+1 in the Rollout/Deployment manifest
2. Argo CD sees the changes in Git and updates the live state in the cluster with the new Rollout object
3. Argo Rollouts takes over as it watches for all changes in Rollout Objects. Argo Rollouts is completely oblivious to what is happening in Git. It only cares about what is happening with Rollout objects that are live in the cluster.
4. Argo Rollouts tries to apply version N+1 with the selected strategy (e.g. blue/green)
5. Version N+1 deploys successfully
6. Cluster is running version N+1 and is completely healthy
7. The Rollout is marked as "Synced" both in ArgoCD and Argo Rollouts.
8. Argo CD syncs take no further action as the Rollout object in Git is exactly the same as in the cluster. They both mention version N+1


## ArgoCD + Argo Rollouts Process (Rollback)


1. Version N runs on the cluster as a Rollout (managed by Argo CD). The Git repository is updated with version N+1 in the Rollout/Deployment manifest
2. Argo CD sees the changes in Git and updates the live state in the cluster with the new Rollout object
3. Argo Rollouts takes over as it watches for all changes in Rollout Objects. Argo Rollouts is completely oblivious to what is happening in Git. It only cares about what is happening with Rollout objects that are live in the cluster.
4. Argo Rollouts tries to apply version N+1 with the selected strategy (e.g. blue/green)
5. Version N+1 fails to deploy for some reason
6. Argo Rollouts scales back again (or switches traffic back) to version N in the cluster. No change in Git takes place from Argo Rollouts
7. Cluster is running version N and is completely healthy
8. The Rollout is marked as "Degraded" both in ArgoCD and Argo Rollouts.
9. Argo CD syncs take no further action as the Rollout object in Git is exactly the same as in the cluster. They both mention version N+1

## In Practice

1. Define Application as an argo rollout resource, pod config goes in spec > template as usual.
2. 