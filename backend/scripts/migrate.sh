#!/bin/bash
set -e

# Database migration script

# Load environment variables if .env exists
if [ -f .env ]; then
    export $(cat .env | grep -v '#' | awk '/=/ {print $1}')
fi

# Default values
DB_HOST=${DB_HOST:-localhost}
DB_PORT=${DB_PORT:-5432}
DB_USER=${DB_USER:-urlshortener}
DB_PASSWORD=${DB_PASSWORD:-urlshortener_password}
DB_NAME=${DB_NAME:-urlshortener}
DB_SSLMODE=${DB_SSLMODE:-disable}

# Construct database URL
DATABASE_URL="postgresql://${DB_USER}:${DB_PASSWORD}@${DB_HOST}:${DB_PORT}/${DB_NAME}?sslmode=${DB_SSLMODE}"

# Run migrations
echo "Running database migrations..."
migrate -path migrations -database "${DATABASE_URL}" up

echo "Migrations completed successfully!"
