#!/bin/bash
set -e

echo "Printing environment variables:"
echo "DB USER: ${DBUSER}"
echo "DB NAME: ${DBNAME}"
echo "DB HOST: ${DBHOST}"
echo "POSTGRES USER: ${POSTGRESUSER}"
echo "Done printing environment variables:"

echo "Wait for Postgres to start"
until pg_isready -h ${DBHOST} -p 5432 -U ${POSTGRESUSER}
do
  echo "Waiting for postgres to start at ${DBHOST}..."
  sleep 2;
done
echo "Postgres has started"

echo "Setting up the Postgres user"
POSTGRES="psql --username ${POSTGRESUSER}"

echo "Creating database: ${DBNAME}"
$POSTGRES <<EOSQL
CREATE DATABASE "${DBNAME}" OWNER ${POSTGRESUSER};
EOSQL

echo "Initializing database..."
psql -d ${DBNAME} -a -U${POSTGRESUSER} -f ./scripts/init.sql &
SQL_PID=$!
wait $SQL_PID
echo "Finished initializing database"

echo "Running Go Tests"
go test -v ./... -coverpkg ./... -coverprofile cover.out &
GO_PID=$!
wait $GO_PID
echo "Finished with go tests"