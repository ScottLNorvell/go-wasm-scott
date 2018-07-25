#!/bin/bash
imageName=go-wasm-scott:image
containerName=go-wasm-scott

docker build -t $imageName -f Dockerfile  .

echo Delete old container...
docker rm -f $containerName

echo Run new container...
docker run -d -p 5000:80 --name $containerName $imageName
