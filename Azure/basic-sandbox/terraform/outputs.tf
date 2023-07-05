output "gateway_frontend_ip" {
  value = "http://${azurerm_public_ip.appgw-public-ip.ip_address}"
}

output "secret_identifier" {
  value = azurerm_key_vault_certificate.cert.secret_id
}