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

## In Practice (Builtin K8s Service - Canary)

1. Define Application as an argo rollout resource, pod config goes in spec > template as usual. Store in git repository
2. Create application and service / other required resourses as an argoCD application with autosync enabled to git repo.
3. Update the manifest in git (Usually the image tag but can be any part)
4. ArgoCD will sync and begin the rollout process, with original manifest as 'stable' and new manifest as 'release'
5. The controller will progress through the steps defined in the Rollout's update strategy. If using builtin service, rollouts doesn't adjust the way service routes traffic, just the number of replicas from the stable manifest vs the release manifest. So if 5 replicas, *setWeight: 20* means that 1 replica will be on release, and the other 4 on stable.
6. Watch the rollout progress with: `kubectl argo rollouts get rollout rollout-name --watch`
7. If a pause is set, the release manifest can be promoted to the next step with `kubectl argo rollouts promote rollout-name`. 
8. The rollout can also be rolled back at any time `kubectl argo rollouts abort rollout-name` which will adapt all the replicas to use the stable manifest.
9. When all steps of a rollout are complete, the replicaset will be marked as stable.