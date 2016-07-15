#!/usr/bin/env bash
# ----------------------------------------------------------
# PURPOSE

# This is the test manager for toadserver. It will run the testing
# sequence for toadserver using docker.

# ----------------------------------------------------------
# REQUIREMENTS

# eris installed locally

# ----------------------------------------------------------
# USAGE

# test.sh

# ----------------------------------------------------------
# Set defaults

# Where are the Things?
##XXX TODO harmozine variable naming convention

name=toadserver
base=github.com/eris-ltd/$name
repo=`pwd`
if [ "$CIRCLE_BRANCH" ]
then
  ci=true
  linux=true
elif [ "$TRAVIS_BRANCH" ]
then
  ci=true
  osx=true
elif [ "$APPVEYOR_REPO_BRANCH" ]
then
  ci=true
  win=true
else
  repo=$GOPATH/src/$base
  ci=false
fi

branch=${CIRCLE_BRANCH:=master}
branch=${branch/-/_}
branch=${branch/\//_}

# Other variables
if [[ "$(uname -s)" == "Linux" ]]
then
  uuid=$(cat /proc/sys/kernel/random/uuid | tr -dc 'a-zA-Z0-9' | fold -w 12 | head -n 1)
elif [[ "$(uname -s)" == "Darwin" ]]
then
  uuid=$(uuidgen | tr -dc 'a-zA-Z0-9' | fold -w 12 | head -n 1  | tr '[:upper:]' '[:lower:]')
else
  uuid="62d1486f0fe5"
fi

was_running=0
test_exit=0
chains_dir=$HOME/.eris/chains
chain_name=toadserver-tests-$uuid
name_full="$chain_name"_full_000
chain_dir=$chains_dir/$chain_name

export ERIS_PULL_APPROVE="true"
export ERIS_MIGRATE_APPROVE="true"

# ---------------------------------------------------------------------------
# Needed functionality

ensure_running(){
  if [[ "$(eris services ls -qr | grep $1)" == "$1" ]]
  then
    echo "$1 already started. Not starting."
    was_running=1
  else
    echo "Starting service: $1"
    eris services start $1 1>/dev/null
    early_exit
    sleep 3 # boot time
  fi
}

early_exit(){
  if [ $? -eq 0 ]
  then
    return 0
  fi

  echo "There was an error duing setup; keys were not properly imported. Exiting."
  if [ "$was_running" -eq 0 ]
  then
    if [ "$ci" = true ]
    then
      eris services stop keys
    else
      eris services stop -rx keys
    fi
  fi
  exit 1
}

test_setup(){
  echo "Getting Setup"
  if [ "$ci" = true ]
  then
    eris init --yes --pull-images=true --testing=true 1>/dev/null
  fi
  ensure_running keys

  # make a chain
  eris chains make --account-types=Full:1 $chain_name 1>/dev/null
  PUBKEY=$(cat $chain_dir/accounts.csv | grep $name_full | cut -d ',' -f 1)
  echo -e "Default PubKey =>\t\t\t\t$PUBKEY"
  eris chains new $chain_name --dir $chain_dir 1>/dev/null
  sleep 5 # boot time
  echo "Setup complete"
}

perform_tests(){
  echo
  echo "starting toadserver"
  eris services start toadserver --chain=$chain_name --env "MINTX_PUBKEY=$PUBKEY" --env "MINTX_CHAINID=$chain_name"
  if [ $? -ne 0 ]
  then
    test_exit=1
    return 1
  fi
  
  sleep 5
  ensure_running ipfs

# wake up ipfs
# XXX shouldn't need this!!
  F_CONTENTS_POST="work pls ipfs"
  F_NAME=guh.txt
  F_PATH=$chain_dir/$F_NAME
  echo "$F_CONTENTS_POST" > $F_PATH
  eris files put $F_PATH -d

  #if [ "$ci" = true ]
  #then
    ## need dm ip
    dm_active=$(docker-machine active)
    dm_ip=$(docker-machine ip $dm_active)
    ERIS_IPFS_HOST="$dm_ip"
  #fi
  
  echo "Generating test file"
  echo ""
  FILE_CONTENTS_POST="testing the toadserver"
  FILE_NAME=hungryToad.txt
  FILE_PATH=$chain_dir/$FILE_NAME

  echo "$FILE_CONTENTS_POST" > $FILE_PATH

  echo "Posting to toadserver"
  here=`pwd`
  cd $chain_dir
  toadserver put $FILE_NAME --host=$dm_ip
  cd $here
  #curl --silent -X POST http://${dm_ip}:11113/postfile?fileName=${FILE_NAME} --data-binary "@$FILE_PATH"

  if [ $? -ne 0 ]
  then
    echo "failed posting to toadserver"
    test_exit=1
    return 1
  fi
  sleep 5 # let all the things happen

  # ask toadserver for the file
  echo "Getting from toadserver"
  toadserver get $FILE_NAME --host=$dm_ip # drop $FILENAME in pwd
  FILE_CONTENTS_GET=`cat $FILE_NAME`

  #FILE_CONTENTS_GET=$(curl --silent -X GET http://${dm_ip}:11113/getfile?fileName=${FILE_NAME}) #output directly or use -o to save to file & read
  if [ $? -ne 0 ]
  then
    test_exit=1
    return 1
  fi
  echo

if [[ "$FILE_CONTENTS_POST" != "$FILE_CONTENTS_GET" ]]; then
    echo "Post contented does not match getted content."
    echo  "Got $FILE_CONTENTS_GET, expected $FILE_CONTENTS_POST"
    test_exit=1
    return 1
  fi
}

test_teardown(){
  if [ "$ci" = false ]
  then
    echo
    eris services stop -rxf toadserver 1>/dev/null
    eris chains stop -f $chain_name 1>/dev/null
    eris chains rm -x --file $chain_name 1>/dev/null
    if [ "$was_running" -eq 0 ]
    then
      eris services stop -rx keys &>/dev/null
    fi

    rm -rf $HOME/.eris/scratch/data/toadserver-tests-*
    rm -rf $chain_dir
  else
    eris chains stop -f $chain_name 1>/dev/null
  fi
  echo
  if [ "$test_exit" -eq 0 ]
  then
    echo "Tests complete! Tests are Green. :)"
  else
    echo "Tests complete. Tests are Red. :("
  fi
  cd $start
  exit $test_exit
}

# ---------------------------------------------------------------------------
# Get the things build and dependencies turned on

echo "Hello! I'm the marmot that tests toadserver."
start=`pwd`
cd $repo
echo ""
echo "Building toadserver in a docker container."
set -e
tests/build_tool.sh 1>/dev/null
set +e
if [ $? -ne 0 ]
then
  echo "Could not build toadserver. Debug via by directly running [`pwd`/tests/build_tool.sh]"
  exit 1
fi
echo "Build complete."
echo ""

# ---------------------------------------------------------------------------
# Setup

test_setup

# ---------------------------------------------------------------------------
# Go!

perform_tests

# ---------------------------------------------------------------------------
# Cleaning up

test_teardown

