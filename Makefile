.DEFAULT_GOAL := help

GO111MODULE := on

help: ## Display this help message
	@echo "Please use \`make <target>' where <target> is one of"
	@perl -nle'print $& if m{^[\.a-zA-Z_-]+:.*?## .*$$}' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m  %-25s\033[0m %s\n", $$1, $$2}'

clean: ## Clean up build artifacts
	rm -rf *.zip bootstrap

build: clean ## Build the Go Lambda function
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bootstrap lambda/main.go

zip: build ## Create a zip file for deployment
	zip bootstrap bootstrap

deploy: zip ## Deploy the Lambda function to AWS
	aws lambda update-function-code --function-name go-sample-lambda --zip-file fileb://bootstrap.zip