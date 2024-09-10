#### General Configs ####

variable "region" {}
variable "cluster_name" {}

variable "service_name" {}
variable "service_port" {}
variable "service_cpu" {}
variable "service_memory" {}
variable "service_launch_type" {}
variable "service_task_count" {}
variable "service_hosts" {}

variable "ssm_vpc_id" {}
variable "ssm_listener" {}
variable "ssm_private_subnet_1" {}
variable "ssm_private_subnet_2" {}
variable "ssm_private_subnet_3" {}

variable "environment_variables" {}
variable "capabilities" {}
variable "service_health_check" {
  type = map(any)
}