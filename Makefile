.PHONY: build-local build templ notify-templ-proxy run

-include .env

build-local:
	@go build -o ./bin/main cmd/main/main.go

build:
	@npm run build
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/main cmd/main/main.go

templ:
	@templ generate --watch --proxy=http://localhost:$(APP_PORT) --proxyport=$(TEMPL_PROXY_PORT) --open-browser=false --proxybind="0.0.0.0"

notify-templ-proxy:
	@templ generate --notify-proxy --proxyport=$(TEMPL_PROXY_PORT)

dev:
	@echo "Starting development environment..."
	@make templ & sleep 1
	@air

vtail:
	@tail -f ./logs/app.log | jq -C

tail:
	@tail -f ./logs/app.log | jq -r '"\(.time) \(.level) \(.msg)"'
