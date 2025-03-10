protoc-gen-stubs:	protoc -I ../proto/ \
	--go_out ./internal/pb --go_opt paths=source_relative \
	--go-grpc_out ./internal/pb --go-grpc_opt paths=source_relative \
	../proto/v1/article.proto

protoc-gen-gw-stubs:	
	protoc -I ../proto --grpc-gateway_out ./internal/pb \
    --grpc-gateway_opt paths=source_relative \
    --grpc-gateway_opt generate_unbound_methods=true \
    ../proto/v1/article.proto
