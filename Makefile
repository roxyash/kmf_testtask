.PHONY: run drun dcrun swag

run:
	go run proxy_service/cmd/main.go

drun:
	docker build -t proxy_service build/proxy_service.Dockerfile

dcrun:
	docker-compose up

swag:
	swag init -d ./proxy_service -g ./cmd/main.go --output proxy_service/docs