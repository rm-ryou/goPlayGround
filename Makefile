.PHONY: default
default:
	go run main.go
	convert dist/result.ppm dist/result.png

.PHONY: test
test:
	go test -v ./...

