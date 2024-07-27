.PHONY: default
default: run convert

.PHONY: run
run:
	go run main.go

.PHONY: test
test:
	go test -v ./...

.PHONY: convert
# convert artifact
convert:
	convert dist/result.ppm dist/result.png
