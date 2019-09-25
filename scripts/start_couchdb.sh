#!/bin/sh

ADMIN_USER=admin
ADMIN_PASSWORD=admin

docker stop local-couchdb > /dev/null 2>&1
docker rm local-couchdb > /dev/null 2>&1
docker pull couchdb:2.3.1
docker run --env COUCHDB_USER=$ADMIN_USER --env COUCHDB_PASSWORD=$ADMIN_PASSWORD \
 -p 4369:4369 -p 5984:5984 -p 9100:9100 \
 -d --name local-couchdb couchdb:2.3.1
