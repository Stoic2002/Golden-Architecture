#!/bin/bash

# Database migration script
# Usage: ./scripts/migrate.sh [up|down]

set -e

DB_HOST="${DB_HOST:-localhost}"
DB_PORT="${DB_PORT:-5432}"
DB_USER="${DB_USER:-arulkarim}"
DB_PASSWORD="${DB_PASSWORD:-root}"
DB_NAME="${DB_NAME:-todo_db}"

MIGRATIONS_DIR="./migrations"

case "$1" in
    up)
        echo "Running migrations up..."
        for file in "$MIGRATIONS_DIR"/*.up.sql; do
            if [ -f "$file" ]; then
                echo "Applying: $file"
                PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f "$file"
            fi
        done
        echo "Migrations completed successfully!"
        ;;
    down)
        echo "Running migrations down..."
        for file in $(ls -r "$MIGRATIONS_DIR"/*.down.sql 2>/dev/null); do
            if [ -f "$file" ]; then
                echo "Reverting: $file"
                PGPASSWORD=$DB_PASSWORD psql -h $DB_HOST -p $DB_PORT -U $DB_USER -d $DB_NAME -f "$file"
            fi
        done
        echo "Rollback completed successfully!"
        ;;
    *)
        echo "Usage: $0 [up|down]"
        exit 1
        ;;
esac
