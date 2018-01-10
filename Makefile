setup:
	@echo 'Installing goimports...'
	@go get -u golang.org/x/tools/cmd/goimports
	@echo 'Installing golint...'
	@go get -u github.com/golang/lint/golint
	@echo 'Installing dep...'
	@go get -u github.com/golang/dep/cmd/dep
	@echo 'Installing test dependencies...'
	@go get -t ./...
	@echo 'Running dep ensure...'
	@dep ensure -v

test:
	@go test --cover ./...

lint:
	@golint ./...

vet:
	@go vet ./...

.DEFAULT_GOAL := setup
