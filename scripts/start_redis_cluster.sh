#!/usr/bin/env bash

# start redis cluster servers
docker run -d --name redis-7001 --net=host redis:latest redis-server --port 7001 --cluster-enabled yes --appendonly yes
docker run -d --name redis-7002 --net=host redis:latest redis-server --port 7002 --cluster-enabled yes --appendonly yes
docker run -d --name redis-7003 --net=host redis:latest redis-server --port 7003 --cluster-enabled yes --appendonly yes

sleep 3s

# create cluster
HOST_IP=$(hostname -I | cut -d' ' -f1)
docker run --rm redis:latest redis-cli --cluster-yes --cluster create $HOST_IP:7001 $HOST_IP:7002 $HOST_IP:7003
