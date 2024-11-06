.DEFAULT_GOAL := bindings
.PHONY: api clean cron references bindings

clean:
	rm -rf build bindings

api:
	make env-decrypt
	go run cmd/api/main.go

cron:
	make env-decrypt
	go run cmd/cron/main.go

references:
	make env-decrypt
	mkdir -p abis
	go run cmd/references/main.go

bindings:
	make env-decrypt
	make clean
	make references
	mkdir -p bindings
	go run cmd/bindings/main.go

env-encrypt:
	go run cmd/env/main.go encrypt

env-decrypt:
	go run cmd/env/main.go decrypt
