#!/bin/sh
##############################
#####Setting Environments#####
echo "Setting Environments"
set -xe
export PATH=$PATH:$GOPATH/bin:$HOME/bin:$GOROOT/bin
##############################
######Install Dependence######
echo "Installing Dependence"
# go get github.com/go-sql-driver/mysql
##############################
#########Running Test#########
echo "Running Test"
mkdir -p build
pkgs=$(go list ./... | grep -v /vendor/)
echo "mode: set" > build/all.cov
for p in $pkgs;
do
 if [ "$1" = "-u" ];then
  go get -u $p
 fi
 go test -v -timeout 20m -covermode count --coverprofile=build/c.cov $p
 cat build/c.cov | grep -v "mode" >> build/all.cov
done

if [ "`uname`" = "Darwin" ];
then
   gocov convert build/all.cov > build/coverage.json
   cat build/all.cov | sed 's/github.com\/CodeSerenade\/ysepaysdk\///' > build/coverage.cov
   cat build/coverage.json | gocov-html > build/coverage.html
   cat build/coverage.cov | gocover-cobertura > build/coverage.xml
   go tool cover -func build/all.cov | grep total
fi
