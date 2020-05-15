#!/bin/bash

tag='latest'

docker build -t go-rest-docker:$tag .
docker run -p 10000:10000 go-rest-docker:$tag -d