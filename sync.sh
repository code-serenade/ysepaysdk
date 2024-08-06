#!/bin/sh
set -xe

cd ~/go/src/github.com/CodeSerenade/easycrypto
git pull
easycrypto_sha=`git rev-parse HEAD`

cd ~/go/src/github.com/CodeSerenade/ysepaysdk
go get github.com/CodeSerenade/easycrypto@$easycrypto_sha

go mod tidy
