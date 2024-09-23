.DEFAULT_GOAL := bindings
.PHONY: clean bindings

clean:
	rm -rf build abigenBindings 

bindings:
	make clean
	pnpm truffle compile
	mkdir -p abigenBindings
	pnpm truffle run abigen
	mkdir -p bindings
	go run cmd/bindings/main.go
