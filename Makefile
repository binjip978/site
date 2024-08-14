.PHONY:
	build docker docker-run

build:
	GOOS=linux GOARCH=amd64 go build

docker:
	docker build --platform linux/amd64 . -t site:latest

docker-run:
	docker run -p 443:443 --restart=always -d site:latest

