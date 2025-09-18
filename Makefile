run:
	go run ./cmd/rest/

swag:
	swag init -g cmd/rest/main.go -o docs --parseDependency --parseInternal

seed:
	go run ./cmd/seeder/

k8s-create-config-map:
	kubectl create configmap mallow-sale-api-env --from-file=./env/.env.local -n mallow-sale

jenkins-up:
	docker compose -f docker-compose.jenkins.yml up -d

jenkins-down:
	docker compose -f docker-compose.jenkins.yml down