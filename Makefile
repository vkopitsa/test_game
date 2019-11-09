# all: run

# build: 
# 	go build cmd/server/main.go

# run: 
# 	go run cmd/server/main.go

RELNAME=game
GOARCH=amd64

# Setup the -ldflags option for go build here, interpolate the variable values
LDFLAGS = -ldflags "-s -w"

all: fmt test build

test: ## test
	go test ./...

build: linux darwin windows ## Build server

server-start: ## start 
	go run cmd/server/main.go

front-build: ## build front
	cd front/ && NODE_ENV=production npm run build

front-dev: ## develop front
	cd front/ && npm start

linux:
	CGO_ENABLED=0 GOOS=linux GOARCH=${GOARCH} go build ${LDFLAGS} -o ./bin/${RELNAME}-linux-${GOARCH} cmd/server/*.go

darwin:
	CGO_ENABLED=0 GOOS=darwin GOARCH=${GOARCH} go build ${LDFLAGS} -o ./bin/${RELNAME}-darwin-${GOARCH} cmd/server/*.go

windows:
	CGO_ENABLED=0 GOOS=windows GOARCH=${GOARCH} go build ${LDFLAGS} -o ./bin/${RELNAME}-windows-${GOARCH}.exe cmd/server/*.go

up: ## To up all containers
	cd docker/ && docker-compose up -d

down: ## To down all containers
	cd docker/ && docker-compose stop

deploy: ## Deploy docker container
	cd docker/ && docker-compose up -d --no-deps --build app && docker-compose restart app

fmt:
	go fmt ./...

# Absolutely awesome: http://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

.DEFAULT_GOAL := help