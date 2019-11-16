include ~/.octopus
export $(shell sed 's/=.*//' ~/.octopus)

export TF_VAR_octo_url=$(OCTO_URL)
export TF_VAR_octo_api_key=$(OCTO_API_KEY)

NAME=terraform-provider-octopusdeploy

build:
	go build -o $(NAME) .

check: build
	terraform init
	terraform plan