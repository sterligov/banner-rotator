#!/bin/bash

cd deployments
#docker-compose -f docker-compose.yml up -d
sleep 5s #wait services like database
docker-compose -f docker-compose.test.yml up --build
EXIT_CODE=$?
#docker-compose -f docker-compose.yml stop
echo ${EXIT_CODE}
exit ${EXIT_CODE}
