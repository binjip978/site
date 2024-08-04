.PHONY:
	build docker

build:
	GOOS=linux GOARCH=amd64 go build

docker:
	docker build --platform linux/amd64 . -t site:latest

