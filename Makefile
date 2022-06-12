SHELL:=/bin/bash

gqlgen:
	cd internal/delivery/graphqlsvc/schema && \
	go get github.com/99designs/gqlgen@v0.17.2 && \
	go run github.com/99designs/gqlgen generate

check-modd-exists:
	@modd --version > /dev/null

run: check-modd-exists
	@modd -f ./.modd/server.modd.conf