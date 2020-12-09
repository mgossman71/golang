#!/bin/bash
version=v1

docker build . -t mgossman71/k8srelay:$version
docker build . -t mgossman71/k8srelay:latest

docker push mgossman71/k8srelay:$version
docker push mgossman71/k8srelay:latest
