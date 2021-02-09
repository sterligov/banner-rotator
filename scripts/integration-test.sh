#!/bin/sh

docker-compose --env-file deployments/.env -f deployments/docker-compose.yml up -d --build
sleep 15s
docker-compose -f deployments/docker-compose.test.yml up --build
EXIT_CODE=$?

docker-compose -f deployments/docker-compose.test.yml down --remove-orphans
docker-compose --env-file deployments/.env -f deployments/docker-compose.yml down --remove-orphans
exit ${EXIT_CODE}
