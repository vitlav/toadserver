#!/usr/bin/env bash
# ----------------------------------------------------------
# PURPOSE

# This is the build script for toadserver. It will build the tool
# into docker containers in a reliable and predicatable
# manner.

# ----------------------------------------------------------
# REQUIREMENTS

# docker installed locally

# ----------------------------------------------------------
# USAGE

# build_tool.sh

# ----------------------------------------------------------
# Set defaults

if [ "$CIRCLE_BRANCH" ]
then
  repo=`pwd`
else
  repo=$GOPATH/src/github.com/eris-ltd/toadserver
fi
branch=${CIRCLE_BRANCH:=master}
branch=${branch/-/_}
testimage=${testimage:="quay.io/eris/toadserver"}

release_min=$(cat $repo/version/version.go | tail -n 1 | cut -d \  -f 4 | tr -d '"')
release_maj=$(echo $release_min | cut -d . -f 1-2)



# ---------------------------------------------------------------------------
# Go!
cd $repo
if [[ "$branch" = "master" ]] 
then
  docker build -t $testimage:latest .
  docker tag $testimage:latest $testimage:$release_maj
  docker tag $testimage:latest $testimage:$release_min
else
  docker build -t $testimage:$release_min .
fi

