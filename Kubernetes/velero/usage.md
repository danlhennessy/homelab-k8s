Current storage location: Azure - velero-backups

# Create Backup
velero backup create all-namespaces-backup

velero backup create default-namespace-backup --include-namespaces default

velero backup create workload-namespaces-backup --exclude-namespaces kube-system

velero backup logs default-namespace-backup

# Schedule Backup
velero create schedule NAME --schedule="0 */6 * * *"

# Describe backup
velero backup describe default-namespace-backup

# Restore from Backup
velero restore create default-namespace-restore --from-backup default-namespace-backup

velero restore describe default-namespace-restore

velero restore logs default-namespace-restore