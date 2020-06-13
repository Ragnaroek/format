test:
	go test ./pkg/...

playground:
	GOOS=js GOARCH=wasm go build -o playground/format.wasm cmd/playground/main.go
	
run-playground: playground
	goexec 'http.ListenAndServe(":8080", http.FileServer(http.Dir("./playground")))'
	
.PHONY: playground run-playground