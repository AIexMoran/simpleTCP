include .env
export

default: server-build server-run

server-run:
	@docker run -d --rm -e SERVER_PORT=$(PORT) -p $(PORT):$(PORT) --name web-server server

server-build:
	@docker build -t server .

.PHONY: server-run server-build default
