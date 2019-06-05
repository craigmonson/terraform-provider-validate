.PHONY: help          # print out help
help:  
	@cat Makefile | grep PHONY | grep -v Makefile

.PHONY: test-cover
test-cover:
	go test ./...

.PHONY: test          # run 'go test'
test:
	cd validate && go test

.PHONY: build         # build the provider as terraform-provider-validate
build:
	go build -o terraform-provider-validate

.PHONY: build-darwin  # go build darwin amd64 and 386 versions
build-darwin:
	GOOS=darwin GOARCH=amd64 go build -o terraform-provider-validate
	tar -zcvf terraform-provider-validate.darwin-amd64.tar.gz terraform-provider-validate
	GOOS=darwin GOARCH=386 go build -o terraform-provider-validate
	tar -zcvf terraform-provider-validate.darwin-386.tar.gz terraform-provider-validate

.PHONY: build-linux   # go build linux amd64 and 386 versions
build-linux:
	GOOS=linux GOARCH=amd64 go build -o terraform-provider-validate
	tar -zcvf terraform-provider-validate.linux-amd64.tar.gz terraform-provider-validate
	GOOS=linux GOARCH=386 go build -o terraform-provider-validate
	tar -zcvf terraform-provider-validate.linux-386.tar.gz terraform-provider-validate

.PHONY: build-doz     # go build windows amd64 and 386 versions
build-doz:
	GOOS=windows GOARCH=amd64 go build -o terraform-provider-validate
	tar -zcvf terraform-provider-validate.windows-amd64.tar.gz terraform-provider-validate
	GOOS=windows GOARCH=386 go build -o terraform-provider-validate
	tar -zcvf terraform-provider-validate.windows-386.tar.gz terraform-provider-validate

.PHONY: build-all     # run build-darwin build-linux and build-doz
build-all: build-darwin build-linux build-doz

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

.PHONY: plan-fail     # terraform plan for FAILING integration test
plan-fail: 
	@echo "THIS SHOULD FAIL TO PLAN"
	@echo "THIS SHOULD FAIL TO PLAN"
	@echo "THIS SHOULD FAIL TO PLAN"
	mv test-provider-pass.tf test-provider-pass.tf.bak
	mv test-provider-fail.tf.bak test-provider-fail.tf
	-terraform plan
	mv test-provider-pass.tf.bak test-provider-pass.tf
	mv test-provider-fail.tf test-provider-fail.tf.bak

.PHONY: tf-test-fail  # terraform init, plan for FAILING integration test
tf-test-fail: init plan-fail
