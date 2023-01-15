#!/bin/bash

echo "%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%"
echo "Seeding the database..."
echo "%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%%"

# Wait for the PostgreSQL server to start
sleep 10

# Connect to the PostgreSQL server and seed the database
psql -h db -U user -W -d dbname -f ./app/initdb.sql