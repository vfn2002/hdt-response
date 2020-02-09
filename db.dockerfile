FROM postgres
COPY .scripts/structure.sql /docker-entrypoint-initdb.d/