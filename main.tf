variable "octo_url" {
  type        = string
  description = "The Octopus URL"
}

variable "octo_api_key" {
  type        = string
  description = "The Octopus API secret key"
}

provider "octopusdeploy" {
  url     = var.octo_url
  api_key = var.octo_api_key
}