.PHONY: gen
gen:
	@protoc -I proto \
		--go_out proto \
		--go_opt paths=source_relative \
		--go-grpc_out proto \
		--go-grpc_opt paths=source_relative \
		proto/helloworld.proto

.PHONY: dev-server
dev-server:
	@go run ./server/main.go

.PHONY: grpcurl
grpcurl:
	@grpcurl \
		-plaintext \
		-import-path proto \
		-proto helloworld.proto \
		localhost:50051 helloworld.Greeter/SayHello
