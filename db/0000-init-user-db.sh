#!/bin/bash
set -e

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
    CREATE DATABASE tinylibrary;
    CREATE USER tinylibrary WITH PASSWORD 'password';
    GRANT ALL PRIVILEGES ON DATABASE tinylibrary to tinylibrary;
EOSQL