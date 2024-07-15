deploy:
	skaffold run

fmt:
	go fmt ./...

run:
	go run main.go

swag:
	docker run --rm -v .:/code ghcr.io/swaggo/swag:latest init
