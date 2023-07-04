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

variable "backend_address_pool_name" {
    default = "myBackendPool"
}

variable "frontend_port_name" {
    default = "myFrontendPort"
}

variable "frontend_ip_configuration_name" {
    default = "myAGIPConfig"
}

variable "http_setting_name" {
    default = "myHTTPsetting"
}

variable "listener_name" {
    default = "myListener"
}

variable "request_routing_rule_name" {
    default = "myRoutingRule"
}

variable "temporary_name_for_rotation" {
  type        = string
  default     = null
  description = "(Optional) Specifies the name of the temporary node pool used to cycle the default node pool for VM resizing. the `var.agents_size` is no longer ForceNew and can be resized by specifying `temporary_name_for_rotation`"
}