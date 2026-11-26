set -e
cd "$(dirname "$0")/.." || exit

mkdir -p ./account/internal/pb/account_api
mkdir -p ./account/internal/pb/swagger
# grpc
protoc -I ./api \
  --go_out=./account/internal/pb --go_opt=paths=source_relative \
  --go-grpc_out=./account/internal/pb --go-grpc_opt=paths=source_relative \
  ./api/account_api/account.proto

# grpc-gateway
protoc -I ./api \
  --grpc-gateway_out=./account/internal/pb --grpc-gateway_opt=paths=source_relative \
  ./api/account_api/account.proto

# swagger
protoc -I ./api \
  --openapiv2_out=./account/internal/pb/swagger \
  ./api/account_api/account.proto
