#!/bin/bash
version=v1.5

docker build . -t mgossman71/k8srelay:$version
docker build . -t mgossman71/k8srelay:latest

docker push mgossman71/k8srelay:$version
docker push mgossman71/k8srelay:latest

#kubectl create deployment -n api k8srelay --image=mgossman71/k8srelay
#kubectl expose deployment --port 8080 -n api k8srelay --type=NodePort --name=k8srelay-svc
