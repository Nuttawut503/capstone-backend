#!/bin/bash

docker run -d \
  --name db \
  --env-file ./.env \
  -e PGDATA=/var/lib/postgresql/data/pgdata \
  -v "$(pwd)"/_data:/var/lib/postgresql/data \
  -p 5432:5432 \
  postgres

docker run -d \
  --name rdb \
  -p 6379:6379 \
  redis
