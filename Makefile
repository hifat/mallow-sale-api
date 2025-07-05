run:
	go run ./cmd/rest/

swag:
	swag init -g cmd/rest/main.go -o docs --parseDependency --parseInternal

seed:
	go run ./cmd/seeder/