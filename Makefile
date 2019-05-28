.PHONY: help    # print out help
help:  
	@cat Makefile | grep PHONY | grep -v Makefile

.PHONY: test    # run 'go test'
test:
	cd validate && go test

.PHONY: build   # build the provider as terraform-provider-validate
build:
	go build -o terraform-provider-validate

.PHONY: init    # terraform init to test integration
init:
	terraform init

.PHONY: plan    # terraform plan to test integration
plan:
	terraform plan

.PHONY: apply   # terraform apply to test integration
apply:
	terraform apply -auto-approve

.PHONY: tf-test # terraform init, plan and apply for integration test
tf-test: init plan apply
	
