#!/bin/bash

docker-compose -f deployments/docker-compose.yml
sleep 10s
docker-compose -f docker/docker-compose.test.yml
exit $$?