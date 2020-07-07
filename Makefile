test:
	go test ./pkg/...

bench:
	go test ./benchmarks/... -bench=.

bench-profile:
	go test ./benchmarks/... -bench=BenchmarkFormatFloat -cpuprofile=cprof
	go tool pprof --text cprof > cprof.txt
	
playground:
	GOOS=js GOARCH=wasm go build -o playground/format.wasm cmd/playground/main.go
	
run-playground: playground
	go run cmd/playground-test/main.go
	
.PHONY: playground run-playground