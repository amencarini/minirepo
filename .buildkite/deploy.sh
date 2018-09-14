#!/bin/bash
for dir in $( ls -d service.* ); do
  cd $dir
  binary=$(echo $dir | rev | cut -d '.' -f 1 | rev)

  DOCKERFILE_CONTENTS="FROM alpine
  ADD service.$binary /
  CMD [\"/service.$binary\"]"

  echo "$DOCKERFILE_CONTENTS" > Dockerfile

  eval $(minikube docker-env)
  GOOS=linux GOARCH=amd64 go build
  docker build -t $binary .
  kubectl delete deployment $binary
  kubectl run $binary --image=$binary:latest --image-pull-policy=Never
  rm service.$binary
  rm Dockerfile
  cd ..
done
