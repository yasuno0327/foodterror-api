#! /bin/bash

set -e

until mysqladmin ping -h db --silent; do
  echo "Waiting for mysqld to be connectable..."
  sleep 1
done

go build

./sandbox-api

exec "$@"