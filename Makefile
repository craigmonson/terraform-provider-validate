.PHONY: help          # print out help
help:  
	@cat Makefile | grep PHONY | grep -v Makefile

.PHONY: test          # run 'go test'
test:
	cd validate && go test

.PHONY: build         # build the provider as terraform-provider-validate
build:
	go build -o terraform-provider-validate

.PHONY: init          # terraform init to test integration
init:
	terraform init

.PHONY: plan          # terraform plan to test integration
plan:
	terraform plan

.PHONY: apply         # terraform apply to test integration
apply:
	terraform apply -auto-approve

.PHONY: tf-test       # terraform init, plan and apply for integration test
tf-test: init plan apply

.PHONY: plan-fail  # terraform plan for FAILING integration test
plan-fail: 
	@echo "THIS SHOULD FAIL TO PLAN"
	@echo "THIS SHOULD FAIL TO PLAN"
	@echo "THIS SHOULD FAIL TO PLAN"
	cp test-provider-pass.tf test-provider-pass.tf.bak
	cp test-provider-fail.tf.bak test-provider-fail.tf
	-terraform plan
	cp test-provider-pass.tf.bak test-provider-pass.tf
	cp test-provider-fail.tf test-provider-fail.tf.bak

.PHONY: tf-test-fail  # terraform init, plan for FAILING integration test
tf-test-fail: init plan-fail
