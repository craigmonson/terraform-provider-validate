.PHONY: help
help:  ## print out help
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.PHONY: test-cover
test-cover:
	go test ./...

.PHONY: test
test: ## run 'go test'
	cd validate && go test

.PHONY: build         # build the provider as terraform-provider-validate
build:
	go build -o terraform-provider-validate

.PHONY: build-darwin
build-darwin: ## go build darwin amd64 and 386 versions
	GOOS=darwin GOARCH=amd64 go build -o terraform-provider-validate-darwin-amd64
	cp terraform-provider-validate-darwin-amd64 terraform-provider-validate
	tar -zcvf terraform-provider-validate.darwin-amd64.tar.gz terraform-provider-validate
	GOOS=darwin GOARCH=386 go build -o terraform-provider-validate-darwin-386
	cp terraform-provider-validate-darwin-386 terraform-provider-validate
	tar -zcvf terraform-provider-validate.darwin-386.tar.gz terraform-provider-validate

.PHONY: build-linux
build-linux: ## go build linux amd64 and 386 versions
	GOOS=linux GOARCH=amd64 go build -o terraform-provider-validate-linux-amd64
	cp terraform-provider-validate-linux-amd64 terraform-provider-validate
	tar -zcvf terraform-provider-validate.linux-amd64.tar.gz terraform-provider-validate
	GOOS=linux GOARCH=386 go build -o terraform-provider-validate-linux-386
	cp terraform-provider-validate-linux-386 terraform-provider-validate
	tar -zcvf terraform-provider-validate.linux-386.tar.gz terraform-provider-validate

.PHONY: build-doz
build-doz: ## go build windows amd64 and 386 versions
	GOOS=windows GOARCH=amd64 go build -o terraform-provider-validate-windows-amd64.exe
	cp terraform-provider-validate-windows-amd64.exe terraform-provider-validate.exe
	tar -zcvf terraform-provider-validate.windows-amd64.tar.gz terraform-provider-validate.exe
	GOOS=windows GOARCH=386 go build -o terraform-provider-validate-windows-386.exe
	cp terraform-provider-validate-windows-386.exe terraform-provider-validate.exe
	tar -zcvf terraform-provider-validate.windows-386.tar.gz terraform-provider-validate.exe

.PHONY: build-all
build-all: build-darwin build-linux build-doz ## run build-darwin build-linux and build-doz

.PHONY: init
init: ## terraform init to test integration
	terraform init

.PHONY: plan
plan: ## terraform plan to test integration
	terraform plan

.PHONY: apply
apply: ## terraform apply to test integration
	terraform apply -auto-approve

.PHONY: tf-test
tf-test: init plan apply ## terraform init, plan and apply for integration test

.PHONY: plan-fail
plan-fail: ## terraform plan for FAILING integration test
	@echo "THIS SHOULD FAIL TO PLAN"
	@echo "THIS SHOULD FAIL TO PLAN"
	@echo "THIS SHOULD FAIL TO PLAN"
	mv test-provider-pass.tf test-provider-pass.tf.bak
	mv test-provider-fail.tf.bak test-provider-fail.tf
	-terraform plan
	mv test-provider-pass.tf.bak test-provider-pass.tf
	mv test-provider-fail.tf test-provider-fail.tf.bak

.PHONY: tf-test-fail
tf-test-fail: init plan-fail ## terraform init, plan for FAILING integration test
