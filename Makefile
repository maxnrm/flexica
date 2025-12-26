.DEFAULT_GOAL=wgo

include .env
export $(shell sed 's/=.*//' .env)

.PHONY:
build:
	docker build . -t flexica

.PHONY: docker-run
docker-run:
	docker run -d --rm --env-file=.env --name flexica flexica:latest

.PHONY: wgo
wgo:
	wgo run main.go