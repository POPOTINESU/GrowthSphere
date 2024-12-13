#!/bin/sh

echo "Starting entrypoint.sh script..."
echo "Database environment variable: $DATABASE"
if [ "$DATABASE" = "postgres" ]
then
    echo "Waiting for postgres..."

    while ! nc -z $DB_HOST 5432; do
      sleep 5
    done


    echo "build project"
    go build -o /app/server ./cmd/main.go

    echo "PostgreSQL started"
fi

exec "$@"