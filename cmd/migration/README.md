# Database Migration

This directory contains the database migration scripts for creating MongoDB collections and indexes.

## Overview

The migration command creates all necessary MongoDB collections and their indexes based on the entities defined in the `internal` modules.

## Collections Created

The migration creates the following collections:

1. **usage_units** - Usage unit reference data
   - Indexes: `code` (unique)

2. **settings** - Application settings

3. **recipe_types** - Recipe type reference data
   - Indexes: `code` (unique), `created_at`

4. **suppliers** - Supplier information
   - Indexes: `name` (unique), `created_at`

5. **users** - User accounts
   - Indexes: `username` (unique), `created_at`

6. **inventories** - Inventory items
   - Indexes: `name`, `created_at`

7. **recipes** - Recipe data
   - Indexes: `name`, `recipe_type.code`, `order_no`, `created_at`

8. **stocks** - Stock records
   - Indexes: `inventory_id`, `supplier_id`, `created_at`

9. **shoppings** - Shopping lists
   - Indexes: `supplier_id`, `status.code`, `created_at`

10. **shopping_inventories** - Shopping inventory items
    - Indexes: `inventory_id`, `supplier_id`, `created_at`

11. **promotions** - Promotion data
    - Indexes: `type.code`, `name`, `created_at`

## Usage

Run the migration using the Makefile command:

```bash
make migrate
```

Or run directly with Go:

```bash
go run ./cmd/migration/
```

## Configuration

The migration uses the same configuration as the main application, loading settings from `./env/.env`.

Make sure your MongoDB connection is properly configured in the environment file before running the migration.

## Behavior

- If a collection already exists, the migration will skip its creation and log a message
- Indexes are created for each collection to optimize query performance
- The migration is idempotent - it can be run multiple times safely
- Total execution time is limited to 60 seconds via context timeout

## Notes

- The migration creates collections and indexes but does NOT seed data
- To seed initial data, use the seeder command: `make seed`
- Run migrations before running the seeder for the first time
