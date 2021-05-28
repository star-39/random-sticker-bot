#!/bin/bash

go build
chmod +x random-sticker-bot

GITHUB_TOKEN=$1

echo "Building container for Github Container Registry!"

c1=$(buildah from fedora-minimal:34)

buildah copy $c1 random-sticker-bot /random-sticker-bot

buildah config --cmd '/random-sticker-bot' $c1

buildah commit $c1 random-sticker-bot

buildah login -u star-39 -p $GITHUB_TOKEN ghcr.io

buildah push random-sticker-bot ghcr.io/star-39/random-sticker-bot:latest
