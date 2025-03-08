protoc-user:
	protoc -I ./proto \
    --go_out ./user-service/internal/pb/ --go_opt paths=source_relative \
	--go-grpc_out ./user-service/internal/pb/ --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./user-service/internal/pb --grpc-gateway_opt paths=source_relative \
	./proto/user.proto

protoc-article:
	protoc -I ./proto \
    --go_out ./article-service/internal/pb/ --go_opt paths=source_relative \
	--go-grpc_out ./article-service/internal/pb/ --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./article-service/internal/pb --grpc-gateway_opt paths=source_relative \
	./proto/article.proto

run-user_sevice:
	cd user-service && go run cmd/server/main.go

run-article_sevice:
	cd article_sevice && go run cmd/server/main.go