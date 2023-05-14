#!/bin/bash
docker run -d --name mysql -e MYSQL_ROOT_PASSWORD=root -e MYSQL_DATABASE=test -p 3306:3306 -v mysql:/var/lib/mysql mysql:latest
docker run -d --name adminer --link mysql -p 8888:8080 adminer:latest
