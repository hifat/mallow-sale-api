run-api:
	go run ./cmd/api ./env/$e/.env.$s

run-all:
	go run ./cmd/api/main.go ./env/local/.env.inventory & \
	go run ./cmd/api/main.go ./env/local/.env.recipe

# Kill all running Go processes
kill-all:
	pkill -f "go run ./cmd"