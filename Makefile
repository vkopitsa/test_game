all: run

build: 
	go build cmd/server/main.go

run: 
	go run cmd/server/main.go