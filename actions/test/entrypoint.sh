#!/usr/bin/env bash

set -e +o pipefail

APP_DIR="/go/src/github.com/${GITHUB_REPOSITORY}/"

mkdir -p ${APP_DIR} && cp -r ./ ${APP_DIR} && cd ${APP_DIR}

go test $(go list ./... | grep -v /vendor/) -race -coverprofile=coverage.txt -covermode=atomic
mv coverage.txt ${GITHUB_WORKSPACE}/coverage.txt
