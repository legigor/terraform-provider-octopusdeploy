include ~/.octopus
export $(shell sed 's/=.*//' ~/.octopus)

export TF_VAR_octo_url=$(OCTO_URL)
export TF_VAR_octo_api_key=$(OCTO_API_KEY)

NAME=terraform-provider-octopusdeploy

clean:
	rm -f $(NAME)
	rm -rf .terraform
	rm -f *.tfstate

build:
	go build -o $(NAME) .


init:
	terraform init

check: clean build init
	terraform plan

apply: clean build init
	terraform apply
