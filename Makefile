run-api:
	go run ./cmd/api ./env/$e/.env.$s

pb-gen:
	protoc --go_out=. --go_opt=paths=source_relative \
    --go-grpc_out=. --go-grpc_opt=paths=source_relative \
    ./internal/$s/$sProto/$s.proto

seed:
	go run ./cmd/seed ./env/$e/.env.$s $s

db-up:
	docker compose -f docker-compose.db.yaml up -d

db-down:
	docker compose -f docker-compose.db.yaml down


docker-run-local:
	docker run --env-file ./env/local/.env.$s -v $(pwd)/env/local/.env.$s:/app/.env --name mallow-sale-api -p $p mallow-sale-api /app/.env
