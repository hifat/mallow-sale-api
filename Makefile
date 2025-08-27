run:
	go run ./cmd/rest/

swag:
	swag init -g cmd/rest/main.go -o docs --parseDependency --parseInternal

seed:
	go run ./cmd/seeder/

k8s-create-config-map:
	kubectl create configmap mallow-sale-api-env --from-file=./env/.env.local -n mallow-sale