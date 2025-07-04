# mallow-sale-api

RESTful API สำหรับจัดการ inventories, recipes, และ recipe_ingredients (ingredients ดึงข้อมูลจาก inventories)

## Tech Stack
- Go
- Gin
- MongoDB
- Clean Architecture

## Features
- CRUD Inventories
- CRUD Recipes
- CRUD Recipe Ingredients (ingredients ดึงข้อมูลจาก inventories)

## Project Structure (Clean Architecture)

```
mallow-sale-api/
  ├── cmd/                # main application entrypoint
  ├── internal/
  │     ├── domain/       # entity, repository interface
  │     ├── usecase/      # business logic
  │     ├── repository/   # mongo implementation
  │     └── handler/      # gin handler (controller)
  ├── pkg/                # (optional) shared packages
  ├── go.mod
  └── README.md
``` 