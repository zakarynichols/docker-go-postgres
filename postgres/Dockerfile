# Postgres - Seed db with .sql script on initialization
FROM postgres:alpine
COPY initdb.sql /docker-entrypoint-initdb.d
RUN chmod a+r /docker-entrypoint-initdb.d/*