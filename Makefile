
.PHONY: proto
proto: proto-user proto-post

.PHONY: proto-user
proto-user:
	protoc --proto_path=user-service/proto \
		--go_out=user-service --go_opt=paths=source_relative \
		--go-grpc_out=user-service --go-grpc_opt=paths=source_relative \
		user-service/proto/user.proto

.PHONY: proto-post
proto-post:
	# Generate user.proto di post-service
	protoc --proto_path=post-service/proto \
		--go_out=post-service --go_opt=paths=source_relative \
		--go-grpc_out=post-service --go-grpc_opt=paths=source_relative \
		post-service/proto/user.proto
	# Generate post.proto di post-service
	protoc --proto_path=post-service/proto \
		--go_out=post-service --go_opt=paths=source_relative \
		--go-grpc_out=post-service --go-grpc_opt=paths=source_relative \
		post-service/proto/post.proto

