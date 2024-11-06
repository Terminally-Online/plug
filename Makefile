.DEFAULT_GOAL := bindings
.PHONY: api clean cron references bindings

clean:
	rm -rf build bindings

api:
	go run cmd/api/main.go

cron:
	go run cmd/cron/main.go

references:
	mkdir -p abis
	go run cmd/references/main.go

bindings:
	make clean
	make references
	mkdir -p bindings
	go run cmd/bindings/main.go

env-encrypt:
	go run cmd/env/main.go encrypt

env-decrypt:
	go run cmd/env/main.go decrypt
