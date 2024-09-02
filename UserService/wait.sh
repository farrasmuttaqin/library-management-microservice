#!/bin/sh

# Tunggu hingga database tersedia menggunakan netcat
while ! nc -z $DATABASE_HOST $DATABASE_PORT; do
  echo "Menunggu database tersedia..."
  sleep 2
done

echo "Database tersedia, memulai aplikasi..."
exec "$@"
