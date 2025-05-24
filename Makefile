run-api:
	go run ./cmd/api ./env/$e/.env.$s

pb-gen:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./internal/$s/$sProto/$s.proto

seed:
	go run ./cmd/seed ./env/$e/.env.$s $s