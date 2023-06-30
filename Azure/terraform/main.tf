provider "azurerm" {
  features {}
}

variable "client_id" {
  description = "Service Principal ID"
}
variable "client_secret" {
  description = "Service Principal secret"
}
variable "create_aks_cluster" {
  type    = bool
  default = true
}

data "azurerm_resource_group" "existing" {
  name     = "AKSsandbox"
}

resource "azurerm_kubernetes_cluster" "sandbox" {
  count               = var.create_aks_cluster ? 1 : 0
  name                = "sandboxcluster"
  location            = data.azurerm_resource_group.existing.location
  resource_group_name = data.azurerm_resource_group.existing.name
  dns_prefix          = "sandboxcluster"

  default_node_pool {
    name                = "default"
    vm_size             = "Standard_B2s"
    os_disk_size_gb     = 30
    node_count          = 1
    enable_auto_scaling = false
  }

  service_principal {
    client_id     = var.client_id
    client_secret = var.client_secret
  }

  tags = {
    environment = "sandbox"
  }
}