.DEFAULT_GOAL := bindings
.PHONY: clean references bindings

clean:
	rm -rf build bindings

references:
	mkdir -p abis
	go run cmd/references/main.go

bindings:
	make clean
	make references
	mkdir -p bindings
	go run cmd/bindings/main.go
