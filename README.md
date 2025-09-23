# Mallow Sale API

## Description

Mallow Sale API is a Go-based REST API service for managing recipes and inventory. It provides comprehensive CRUD operations for recipe management with MongoDB as the database backend.

## Tech Stack

- **Language**: Go
- **Framework**: Gin (HTTP web framework)
- **Database**: MongoDB
- **Documentation**: Swagger/OpenAPI
- **Error Handling**: Custom error handling package

## Quick Start

1. Create a `.env` file and modify it and change `DB_HOST` to `host.docker.internal`

```sh
cp ./env/.env.example ./env/.env
```

2. Run with docker

```sh
docker compose up
```
