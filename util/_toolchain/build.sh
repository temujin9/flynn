#!/bin/bash

set -eo pipefail

version=1.8
shasum=53ab94104ee3923e228a2cb2116e5e462ad3ebaeea06ff04463479d7f12d27ca
pkg="go${version}.linux-amd64"
go=go/bin/go

test -f ${go} && ${go} version | grep -q ${version} && exit

rm -rf go.tar.gz go

curl -o go.tar.gz -L "https://storage.googleapis.com/golang/${pkg}.tar.gz"
echo "${shasum}  go.tar.gz" | shasum -c -

tar xzf go.tar.gz
rm go.tar.gz
