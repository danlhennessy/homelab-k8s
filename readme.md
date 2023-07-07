
## Cluster Features
<br> 
<table>
    <tr>
        <th>Category</th>
        <th>Feature</th>
        <th></th>
        <th>Tool</th>
        <th>Description</th>
    </tr>
    <tr>
        <td rowspan="3">Cluster Core</td>
        <td>Container Runtime</td>
        <td><img width="68" src="https://cncf-branding.netlify.app/img/projects/containerd/stacked/black/containerd-stacked-black.png"></td>
        <td><a href="https://containerd.io/">containerd</a></td>
        <td>Image and Container management, CRI tool</td>
    </tr>
    <tr>
        <td>Cluster Initialization</td>
        <td><img width="68" src="https://user-images.githubusercontent.com/22591623/59856252-74656b00-936e-11e9-8dd9-d6092845981b.png"></td>
        <td><a href="https://kubernetes.io/docs/reference/setup-tools/kubeadm/">Kubeadm</a></td>
        <td>Kubernetes cluster setup, initialization and node joining</td>
    </tr>
    <tr>
        <td>Container Network Interface</td>
        <td><img width="68" src="https://avatars.githubusercontent.com/u/78555908?s=280&v=4"></td>
        <td><a href="https://github.com/flannel-io/flannel">Flannel</a></td>
        <td>Pod network add-on, enables cluster DNS</td>
    </tr>
    <tr>
        <td rowspan="4">Deployment Automation</td>
        <td>CI</td>
        <td><img width="68" src="https://avatars.githubusercontent.com/u/44036562?s=200&v=4)"></td>
        <td><a href=https://github.com/features/actions>Github Actions</a></td>
        <td>Continuous Integration - unpacked and packed testing</td>
    </tr>
    <tr>
        <td>CD</td>
        <td><img width="68" src="https://cncf-branding.netlify.app/img/projects/argo/icon/color/argo-icon-color.svg"></td>
        <td><a href="https://argoproj.github.io/cd">ArgoCD + Rollouts</a></td>
        <td>Continuous Delivery and Deployment - Git synchronised with cluster environments. Canary rollouts</td>
    </tr>
    <tr>
        <td>Infrastructure Provisioning</td>
        <td><img width="68" src="https://cncf-branding.netlify.app/img/projects/crossplane/icon/color/crossplane-icon-color.png"></td>
        <td><a href="https://www.crossplane.io/">Crossplane</a></td>
        <td>For CRDs and linked with Azure provider to manage out-of-cluster cloud resources</td>
    </tr>
    <tr>
        <td>Release Management</td>
        <td><img width="68" src="https://cncf-branding.netlify.app/img/projects/helm/icon/color/helm-icon-color.png"></td>
        <td><a href="https://helm.sh/">Helm</a></td>
        <td>Package management - app versioning, repositories and dependencies</td>
    </tr>
    <tr>
        <td rowspan="2">Storage</td>
        <td>Storage Operator</td>
        <td><img width="68" src="https://cncf-branding.netlify.app/img/projects/rook/icon/color/rook-icon-color.png"></td>
        <td><a href="https://rook.io/">Rook</a></td>
        <td>Cloud native storage orchestrator - link with distributed storage solution</td>
    </tr>
    <tr>
        <td>Disaster Recovery, Migration</td>
        <td><img width="68" src="https://velero.io/img/Velero.svg"></td>
        <td><a href="https://velero.io/">Velero</a></td>
        <td>Data backups, migration and disaster recovery for cluster resources and PVs</td>
    </tr>
    <tr>
        <td rowspan="2">Networking</td>
        <td>Service Mesh</td>
        <td><img width="62" src="https://upload.wikimedia.org/wikipedia/commons/thumb/a/a1/Istio-bluelogo-nobackground-unframed.svg/1365px-Istio-bluelogo-nobackground-unframed.svg.png"></td>
        <td><a href="https://istio.io/">Istio</a></td>
        <td>Cluster networking management and observability</td>
    </tr>
    <tr>
        <td>Health and Topology</td>
        <td><img width="62" src="https://s3.amazonaws.com/media-p.slid.es/uploads/671898/images/6136934/kiali_logo_1color_013144_1280px.svg"></td>
        <td><a href="https://kiali.io/">Kiali</a></td>
        <td>Istio add on to monitor service communication and availability</td>
    </tr>
    <tr>
        <td rowspan="3">Security</td>
        <td>Runtime Security</td>
        <td><img width="68" src="https://cncf-branding.netlify.app/img/projects/falco/stacked/color/falco-stacked-color.png"></td>
        <td><a href="https://falco.org/">Falco</a></td>
        <td>Custom behavioral rules and alerts for cluster anomalies</td>
    </tr>
    <tr>
        <td>Vulnerability Scanning</td>
        <td><img width="68" src="https://api.civo.com/k3s-marketplace/kube-hunter.png"></td>
        <td><a href="https://github.com/aquasecurity/kube-hunter">kube-hunter</a></td>
        <td>Dynamic security tests (passive and active). Highlights potential threats </td>
    </tr>
    <tr>
        <td>Secrets Management</td>
        <td><img width="68" src="https://www.nicepng.com/png/full/827-8272881_vault-logo-black-and-white-hashicorp-vault-logo.png"></td>
        <td><a href="https://www.vaultproject.io/">Vault</a></td>
        <td>Hashicorp Vault cluster for all project secrets, connection established with Github Actions environment secrets</td>
    </tr>
    <tr>
        <td rowspan="2">Policy Enforcement</td>
        <td>Policy Engine</td>
        <td><img width="68" src="https://landscape.cncf.io/logos/open-policy-agent-opa.svg"></td>
        <td><a href="https://www.openpolicyagent.org/">Open Policy Agent</a></td>
        <td>Policy-based control - manages and evaluates policies against incoming requests and configuration changes</td>
    </tr>
    <tr>
        <td>Manifest Validation</td>
        <td><img width="68" src="https://assets-global.website-files.com/61c02e339c1199782326e3ce/61c02e339c1199792d26e436_datree__logo.svg"></td>
        <td><a href="https://www.datree.io/">Datree</a></td>
        <td>Scans configs to prevent misconfigured resources reaching cluster</td>
    </tr>
    <tr>
        <td rowspan="4">Observability</td>
        <td>Monitoring</td>
        <td><img width="68" src="https://artifacthub.io/image/0503add5-3fce-4b63-bbf3-b9f649512a86@1x"></td>
        <td><a href="https://artifacthub.io/packages/helm/prometheus-community/kube-prometheus-stack">kube-prometheus-stack</a></td>
        <td>Monitoring, alerting and visualisation stack including Prometheus, Alert Manager and Grafana</td>
    </tr>
    <tr>
        <td>Logging</td>
        <td><img width="68" src="https://seeklogo.com/images/E/elasticsearch-logo-C75C4578EC-seeklogo.com.png"></td>
        <td><a href="https://www.elastic.co/guide/en/cloud-on-k8s/2.8/k8s-quickstart.html">ECK</a></td>
        <td>Log management and delivery with Elastic Cloud on Kubernetes. Beats -> ElasticSearch -> Kibana</td>
    </tr>
    <tr>
        <td>APM</td>
        <td><img width="68" src="https://static-www.elastic.co/v3/assets/bltefdd0b53724fa2ce/blt5ceddec3c8f0ca55/5d082bb6877575d0584761ac/logo-apm-32-color.svg"></td>
        <td><a href="https://www.elastic.co/guide/en/cloud-on-k8s/2.8/k8s-apm-server.html">Elastic APM Server</a></td>
        <td>Application Performance Monitoring - integrated with ECK service</td>
    </tr>
    <tr>
        <td>Troubleshooting</td>
        <td><img width="68" src="https://lh5.googleusercontent.com/-hf9J6_pbnTk/AAAAAAAAAAI/AAAAAAAAAAA/-Ewgawd0NH4/s44-p-k-no-ns-nd/photo.jpg"></td>
        <td><a href="https://komodor.com/">Komodor</a></td>
        <td>Cluster troubleshooting platform - health monitoring and automated root cause detection</td>
    </tr>
    <!-- Add more rows for other tools -->
</table>
<br>