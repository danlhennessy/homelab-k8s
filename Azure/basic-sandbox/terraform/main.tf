provider "azurerm" {
  features {}
}

resource "random_string" "rg" {
  length  = 8
  upper   = false
  special = false
}

resource "azurerm_resource_group" "existing" {
  name     = "101-sandbox-${random_string.rg.result}"
  location = "uksouth"
}

# Networking

resource "azurerm_virtual_network" "spoke-vnet" {
  name                = "spoke-vnet"
  address_space       = ["10.0.0.0/16"]
  location            = azurerm_resource_group.existing.location
  resource_group_name = azurerm_resource_group.existing.name
}

resource "azurerm_subnet" "aks-subnet" {
  name                 = "aks-subnet"
  resource_group_name  = azurerm_resource_group.existing.name
  virtual_network_name = azurerm_virtual_network.spoke-vnet.name
  address_prefixes     = ["10.0.1.0/24"]
}

resource "azurerm_subnet" "appgw-subnet" {
  name                 = "appgw-subnet"
  resource_group_name  = azurerm_resource_group.existing.name
  virtual_network_name = azurerm_virtual_network.spoke-vnet.name
  address_prefixes     = ["10.0.2.0/24"]
}

resource "azurerm_public_ip" "appgw-public-ip" {
  name                = "appgw-public-ip"
  resource_group_name = azurerm_resource_group.existing.name
  location            = azurerm_resource_group.existing.location
  allocation_method   = "Static"
  sku                 = "Standard"
}


resource "azurerm_application_gateway" "main" {
  name                = "myAppGateway"
  resource_group_name = azurerm_resource_group.existing.name
  location            = azurerm_resource_group.existing.location

  sku {
    name     = "WAF_v2"
    tier     = "WAF_v2"
    capacity = 1
  }

  gateway_ip_configuration {
    name      = "my-gateway-ip-configuration"
    subnet_id = azurerm_subnet.appgw-subnet.id
  }

  frontend_port {
    name = "myFrontendPort"
    port = 80
  }

  frontend_ip_configuration {
    name                 = "myAGIPConfig"
    public_ip_address_id = azurerm_public_ip.appgw-public-ip.id
  }

  backend_address_pool {
    name = "myBackendPool"
    fqdns   = ["10.96.0.4"]
  }

  backend_http_settings {
    name                  = "myHTTPsetting"
    cookie_based_affinity = "Disabled"
    port                  = 80
    protocol              = "Http"
    request_timeout       = 60
  }

  http_listener {
    name                           = var.listener_name
    frontend_ip_configuration_name = var.frontend_ip_configuration_name
    frontend_port_name             = var.frontend_port_name
    protocol                       = "Http"
  }

  request_routing_rule {
    name                       = var.request_routing_rule_name
    rule_type                  = "Basic"
    http_listener_name         = var.listener_name
    backend_address_pool_name  = var.backend_address_pool_name
    backend_http_settings_name = var.http_setting_name
    priority                   = 1
  }

  waf_configuration {
    enabled = true
    firewall_mode            = "Detection"
    rule_set_version         = "3.1"
    file_upload_limit_mb     = 100
    max_request_body_size_kb = 128
  }
}

# AKS

resource "azurerm_kubernetes_cluster" "sandboxcluster" {
  count               = var.create_aks_cluster ? 1 : 0
  name                = "sandboxcluster"
  location            = azurerm_resource_group.existing.location
  resource_group_name = azurerm_resource_group.existing.name
  dns_prefix          = "sandboxcluster"

  default_node_pool {
    name                = "default"
    vm_size             = "Standard_B2ms"
    os_disk_size_gb     = 30
    node_count          = 2
    enable_auto_scaling = false
    vnet_subnet_id = azurerm_subnet.aks-subnet.id
    temporary_name_for_rotation  = var.temporary_name_for_rotation
  }

  

  service_principal {
    client_id     = var.client_id
    client_secret = var.client_secret
  }

  network_profile {
    network_plugin = "kubenet"

    service_cidr     = "10.96.0.0/16"
    dns_service_ip   = "10.96.0.10"
  }

  service_mesh_profile {
    mode = "Istio"
  }

  tags = {
    environment = "sandbox"
  }
}