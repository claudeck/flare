#!/bin/bash
VERSION=0.4.3
COMMIT=$(git rev-parse --short HEAD)

echo "最近版本；$VERSION / $COMMIT"
echo "$VERSION">RELEASE_VERSION
echo "$COMMIT">RELEASE_COMMIT

DOCKERHUB_REPO="ct01:5000/flare"

docker build -t "flare-base:$VERSION" -f docker/manual/Dockerfile.base .
docker build -t "$DOCKERHUB_REPO:$VERSION-amd64" --build-arg FLARE_BASE_IMAGE="flare-base:$VERSION" -f docker/manual/Dockerfile.amd64 .
# docker build -t "$DOCKERHUB_REPO:$VERSION-arm32v7" --build-arg FLARE_BASE_IMAGE="flare-base:$VERSION" -f docker/manual/Dockerfile.arm32v7 .
# docker build -t "$DOCKERHUB_REPO:$VERSION-arm64v8" --build-arg FLARE_BASE_IMAGE="flare-base:$VERSION" -f docker/manual/Dockerfile.arm64v8 .

rm RELEASE_VERSION
rm RELEASE_COMMIT

docker images | grep "$DOCKERHUB_REPO"

docker push ct01:5000/flare:$VERSION-amd64