SHELL := /bin/bash

run:
	go run main.go grav.go

compile:
	go build -o bin/
	GOOS=windows GOARCH=amd64 go build -o bin/
	GOOS=darwin GOARCH=amd64 go build -o bin/