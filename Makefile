test:
	go test ./pkg/...

playground:
	GOOS=js GOARCH=wasm go build -o playground/format.wasm cmd/playground/main.go
	
run-playground: playground
	go run cmd/playground-test/main.go
	
.PHONY: playground run-playground