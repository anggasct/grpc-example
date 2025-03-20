# User Service

Service gRPC untuk manajemen user.

## Struktur Database

Table: users
- id (SERIAL PRIMARY KEY)
- name (VARCHAR(255))

## Environment Variables

```
DB_HOST=localhost
DB_PORT=5432
DB_NAME=user_db
DB_USER=postgres
DB_PASSWORD=
SERVER_PORT=50051
```

## Setup

1. Setup database:
```sql
CREATE DATABASE user_db;
```

2. Install dependencies:
```bash
go mod tidy
```

3. Generate proto:
```bash
protoc --go_out=pb/user --go_opt=paths=source_relative \
    --go-grpc_out=pb/user --go-grpc_opt=paths=source_relative \
    -I=proto proto/user.proto
```

4. Run service:
```bash
go run cmd/server/main.go
```

## gRPC API

### GetUser
Request:
```protobuf
message GetUserRequest {
  int64 id = 1;
}
```

Response:
```protobuf
message GetUserResponse {
  User user = 1;
}

message User {
  int64 id = 1;
  string name = 2;
}
