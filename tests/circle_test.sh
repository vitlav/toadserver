#!/usr/bin/env bash
# ----------------------------------------------------------
# PURPOSE

# This is the test manager for toadserver to be ran from circle ci.
# It will run the testing sequence for toadserver using docker.

# ----------------------------------------------------------
# REQUIREMENTS

# docker installed locally
# docker-machine installed locally
# eris installed locally

# ----------------------------------------------------------
# USAGE

# circle_test.sh

# ----------------------------------------------------------
# Set defaults

uuid=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 12 | head -n 1)
machine="eris-test-toadserver-$uuid"
start=`pwd`

# ----------------------------------------------------------
# Get machine sorted

echo "Setting up a Machine for Marmot Testing"
docker-machine create --driver amazonec2 $machine 1>/dev/null
eval $(docker-machine env $machine)
echo "Machine setup."
echo
docker version
echo

# ----------------------------------------------------------
# Run tests

tests/test.sh
test_exit=$?

# ----------------------------------------------------------
# Cleanup

echo
echo
echo "Cleaning up"
docker-machine rm --force $machine
cd $start
exit $test_exit
