.PHONY: proto
proto:
	protoc --micro_out=paths=source_relative:. --go_out=paths=source_relative:. proto/content.proto

.PHONY: build
build:
	go build -o content ./*.go
