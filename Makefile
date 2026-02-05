rest:
	go run ./cmd/rest/ -envPath=./env/.env

grpc:
	go run ./cmd/grpc/ -envPath=./env/.env

swag:
	swag init -g cmd/rest/main.go -o docs --parseDependency --parseInternal

pb-gen:
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		./pkg/grpc/inventoryProto/inventory.proto

seed:
	go run ./cmd/seeder/

migrate:
	go run ./cmd/migration/

k8s-create-config-map:
	kubectl create configmap mallow-sale-api-env --from-file=./env/.env.local -n mallow-sale

infra-up:
	docker compose -f docker-compose.infra.yml up -d

infra-down:
	docker compose -f docker-compose.infra.yml down