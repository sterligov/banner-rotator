#!/bin/sh

docker-compose --env-file deployments/.env -f deployments/docker-compose.yml up -d --build
while [ "$(docker ps | grep "rotator_migrations")" != "" ]; do :; done

docker-compose -f deployments/docker-compose.test.yml up --build
EXIT_CODE=$?

docker-compose -f deployments/docker-compose.test.yml down
docker-compose --env-file deployments/.env -f deployments/docker-compose.yml down
exit ${EXIT_CODE}
