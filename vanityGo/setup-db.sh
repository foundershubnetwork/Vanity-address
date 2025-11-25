#!/bin/bash

echo "Setting up database schema..."
source .env
DB_URL=${DATABASE_URL}
psql "$DB_URL" -f schema.sql
echo "Database setup complete!"
