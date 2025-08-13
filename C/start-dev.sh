#!/bin/bash

echo "Starting backend development server..."
echo ""
echo "Backend will run on: http://localhost:8081"
echo "API endpoints available at: http://localhost:8081/api"
echo ""

cd "$(dirname "$0")/api/user"
go run main.go