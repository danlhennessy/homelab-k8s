# falco

helm repo add falcosecurity https://falcosecurity.github.io/charts
helm repo update

helm install falco --set driver.kind=ebpf --set tty=true falcosecurity/falco --set falcosidekick.enabled=true --set falcosidekick.webui.enabled=true


# falcosidekick

kubectl port-forward svc/falco-falcosidekick-ui --address 0.0.0.0 2802 &> /dev/null &

Generate example events:
    kubectl run event-generator --image falcosecurity/event-generator -- run syscall --loop