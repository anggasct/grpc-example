# Post Service

Service gRPC untuk manajemen post yang terintegrasi dengan User Service.

## Struktur Database

Table: posts
- id (SERIAL PRIMARY KEY)
- user_id (BIGINT)
- content (TEXT)

## Environment Variables

```
DB_HOST=localhost
DB_PORT=5432
DB_NAME=post_db
DB_USER=postgres
DB_PASSWORD=
SERVER_PORT=50052
USER_SERVICE_HOST=localhost
USER_SERVICE_PORT=50051
```

## Setup

1. Setup database:
```sql
CREATE DATABASE post_db;
```

2. Install dependencies:
```bash
go mod tidy
```

3. Generate proto:
```bash
# Generate post proto
protoc --go_out=pb/post --go_opt=paths=source_relative \
    --go-grpc_out=pb/post --go-grpc_opt=paths=source_relative \
    -I=proto proto/post.proto

# Generate user proto (diperlukan untuk integrasi dengan User Service)
protoc --go_out=pb/user --go_opt=paths=source_relative \
    --go-grpc_out=pb/user --go-grpc_opt=paths=source_relative \
    -I=proto proto/user.proto
```

4. Pastikan User Service sudah running di port 50051

5. Run service:
```bash
go run cmd/server/main.go
```

## gRPC API

### GetPost
Request:
```protobuf
message GetPostRequest {
  int64 id = 1;
}
```

Response:
```protobuf
message GetPostResponse {
  Post post = 1;
}

message Post {
  int64 id = 1;
  int64 user_id = 2;
  string content = 3;
  User user = 4;  // Data user dari User Service
}
```

## Integrasi
Service ini terintegrasi dengan User Service untuk mendapatkan data user ketika mengambil data post. Pastikan User Service sudah running sebelum menggunakan endpoint GetPost.
