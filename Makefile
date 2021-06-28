build:
	go mod vendor
	go fmt ./...
	go build -o build/opnsense/opnsense cmd/opnsense/opnsense.go
clean:
	rm -rf build
	go clean
fmt:
	go fmt ./...
