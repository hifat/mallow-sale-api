run-rest:
	go run ./cmd/rest

run-all:
	go run ./cmd/rest/inventory.go ./env/local/.env.inventory & \
	go run ./cmd/rest/recipe.go ./env/local/.env.recipe & \
	go run ./cmd/rest/usageUnit.go ./env/local/.env.usageUnit & \

# Kill all running Go processes
kill-all:
	pkill -f "go run server/"