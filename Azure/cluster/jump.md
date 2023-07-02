# Jumpbox Installs
## AZ CLI
curl -sL https://aka.ms/InstallAzureCLIDeb | sudo bash
## Brew
sudo apt-get install build-essential procps curl file git
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
(echo; echo 'eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"') >> /home/$USER/.bashrc
eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"
## kubelogin
brew install Azure/kubelogin/kubelogin
## kubectl
curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"
sudo install -o root -g root -m 0755 kubectl /usr/local/bin/kubectl
## k9s
brew install derailed/k9s/k9s

# Jumpbox Setups
## Authenticate
az login && az account set --subscription (Subscription ID)
az aks get-credentials --resource-group AKSsandbox --name Portal-AKS  # Replace RG and AKS name as needed
kubelogin convert-kubeconfig -l azurecli