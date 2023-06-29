# Deployment options

Application repos with their own app CI/CD , ending in a new image being pushed to a container registry

Dev, Staging and Prod Namespaces

How to separate dev/stg/prd manifests
    put app manifests in app repo
    In branches so that dont need PRs to test. Dev branch that argo monitors for dev deployments
    Prd branch for staging and prod deployments

# Choices

When to build dev cluster (automatic or not, after app CI/CD or before):
    Optional build at end of CI/CD (Can choose to build at any time during pipeline)
    Also have a docker sandbox for quick non deployment tests
When to build stg cluster (automatic or not, after app CI/CD or before):
    After app CI/CD
    Not automatically
    Stg cluster mirrors prod except for scaling and comes from master branch
