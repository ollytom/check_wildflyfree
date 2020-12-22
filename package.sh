#!/bin/sh

# creates a tarball in /tmp to be uploaded to github releases

GOOS=linux
GOARCH=amd64
version=`git tag |tail -1`

go build 

mkdir -p /tmp/check_wildflyfree
cp check_wildflyfree /tmp/check_wildflyfree
cp check_wildflyfree.1 /tmp/check_wildflyfree
cp README.md /tmp/check_wildflyfree
cp icinga2.conf /tmp/check_wildflyfree
cp LICENSE /tmp/check_wildflyfree

(cd /tmp && tar cv check_wildflyfree | gzip -c > check_wildflyfree-$version-${GOOS}_${GOARCH}.tar.gz)
