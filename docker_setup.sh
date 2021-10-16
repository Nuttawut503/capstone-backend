#!/bin/bash

docker run -d \
  --name db \
  --env-file ./.env \
  -e PGDATA=/var/lib/postgresql/data/pgdata \
  -v "$(pwd)"/data:/var/lib/postgresql/data \
  -p 5432:5432 \
  postgres
