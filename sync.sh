#!/bin/sh
set -xe

cd ~/go/src/github.com/code-serenade/easycrypto
git pull
easycrypto_sha=`git rev-parse HEAD`

cd ~/go/src/github.com/code-serenade/ysepaysdk
go get github.com/code-serenade/easycrypto@$easycrypto_sha

go mod tidy
