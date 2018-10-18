#!/bin/bash

docker container rm -f mariadb 2>/dev/null

docker run --name mariadb -p 3306:3306 -v /docker/mysql:/var/lib/mysql -e MYSQL_ROOT_PASSWORD=s3cr3t -d mariadb:10.2

