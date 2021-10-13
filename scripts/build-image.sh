#!/bin/bash

MODULE=$1
DOCKER_FILE=Dockerfile-edgeUser
TIME=`date "+%Y%m%d%H%M"`
GIT_REVISION=`git describe --tags --always`
IMAGE_NAME=193.168.1.211:8082/microservice/${MODULE}:${TIME}_${GIT_REVISION}
docker build  -t ${IMAGE_NAME} -f ${DOCKER_FILE} .
docker push ${IMAGE_NAME}
echo "${IMAGE_NAME}" > IMAGE_NAME
