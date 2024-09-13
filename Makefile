run:
	go run cmd/$(filter-out $@,$(MAKECMDGOALS))/main.go

proto:
	protoc --go_out=api/ --go_opt=paths=source_relative \
		--go-grpc_out=api/ --go-grpc_opt=paths=source_relative \
		-I api/src api/src/*.proto
