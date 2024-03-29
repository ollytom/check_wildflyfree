#!/bin/sh

# creates a tarball in /tmp to be uploaded to github releases

GOOS=linux
GOARCH=amd64
version=`git tag |tail -1`

go build -o check_wildflyfree

mkdir -p /tmp/check_wildflyfree
cp check_wildflyfree /tmp/check_wildflyfree
cp check_wildflyfree.1 /tmp/check_wildflyfree
cp README.md /tmp/check_wildflyfree
cp icinga2.conf /tmp/check_wildflyfree
cp LICENSE /tmp/check_wildflyfree

tarball="/tmp/check_wildflyfree-$version-${GOOS}_${GOARCH}.tar.gz"
(cd /tmp && tar c check_wildflyfree | gzip -c > $tarball)
echo $tarball
