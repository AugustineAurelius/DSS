
build:
	@go clean -cache
	@go build -o bin/fs cmd/main.go 

run: build
	@./bin/fs


test: cleanup
	@go test ./... -v -benchmem


bench: cleanup
	@go test ./... -bench=. -benchmem

cleanup:
	@go clean -testcache



