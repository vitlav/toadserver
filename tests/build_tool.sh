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

testimage=${testimage:="quay.io/eris/toadserver"}

cd $repo

# ---------------------------------------------------------------------------
# Go!

docker build -t $testimage:latest .
