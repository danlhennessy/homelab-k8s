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