
build:
	@go build -o bin/fs

run: build
	@./bin/fs


test: cleanup
	@go test ./... -v -benchmem


bench: cleanup
	@go test ./... -bench=. -benchmem

cleanup:
	@go clean -testcache

