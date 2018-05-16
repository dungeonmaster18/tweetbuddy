.PHONY: build
build: ## Build the docker Image for this project.
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix nocgo -o /dm-it

.PHONY: run
run: ## Run the app.
	go run main.go

.PHONY: dep
dep: ## Run the app.
	go get -u github.com/golang/dep/cmd/dep
	dep ensure

.PHONY: lint
lint: ## Run the app.
	go get -u golang.org/x/lint/golint
	golint -set_exit_status $(go list ./... | grep -v /vendor/)

.PHONY: help
help: ## Shows help.
	@echo
	@echo 'Usage:'
	@echo '    make <target>'
	@echo
	@echo 'Targets:'
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "    \033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo
