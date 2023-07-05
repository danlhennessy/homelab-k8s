output "gateway_frontend_ip" {
  value = "http://${azurerm_public_ip.appgw-public-ip.ip_address}"
}