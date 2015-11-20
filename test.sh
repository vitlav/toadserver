#! /bin/bash

echo "Starting keys"

eris services start keys
sleep 3

echo "Generating a key"
ADDR=$(eris keys gen)

echo "OUR ADDRESS:"
echo "$ADDR"
echo ""

echo "Exporting keys from container to host"
eris keys export

#cat ~/.eris/keys/data/$ADDR/$ADDR

CHAIN_NAME=toadserver_test

echo "Setting chain name:"
echo "$CHAIN_NAME"
echo ""

CHAIN_DIR=~/.eris/chains/$CHAIN_NAME
mkdir $CHAIN_DIR

echo "Setting and making chain directory:"
echo "$CHAIN_DIR"
echo ""

echo "Converting key to tendermint format:"
PRIV=$(eris keys convert $ADDR)
echo "$PRIV"
echo ""

echo "Piping key to "$CHAIN_DIR"/priv_validator.json"
echo "$PRIV" > "$CHAIN_DIR/priv_validator.json"

echo "Setting pubkey:"
PUB=$(eris keys pub $ADDR)
echo "$PUB"
echo ""

echo "Making genesis file:"
GEN=$(eris chains make-genesis $CHAIN_NAME $PUB)
echo "$GEN"
echo ""

echo "Piping genesis file to "$CHAIN_DIR"/genesis.json"
echo "$GEN" > "$CHAIN_DIR/genesis.json"

echo "Copying default config to "$CHAIN_DIR"/default.toml"
cp ~/.eris/chains/default/config.toml $CHAIN_DIR/
echo ""

echo "Starting chain"
eris chains new $CHAIN_NAME --dir $CHAIN_DIR
#sleep 5
#check that it is running
#eris chains ls --running --quiet
#if...

#ok, chain running, lets boot up the toadserver

echo "Setting service definition file:"

#CHAIN_NAME_1=$CHAIN_NAME"_1" ##dumb hack
CHAIN_NAME_1="${CHAIN_NAME}_1"
echo "$PUB"
#PK="pk=${PUB}"

PK=${PUB//[^A-Z0-9]/}
echo "$PK"

#"MINTX_PUBKEY=$PUB",
read -r -d '' SERV_DEF << EOM
name = "toadserver_test"

[service]
name = "toadserver_test"
image = "quay.io/eris/toadserver"
ports = [ "11113:11113" ]
volumes = [  ]
environment = [  
"MINTX_NODE_ADDR=http://eris_chain_$CHAIN_NAME_1:46657/",
"MINTX_CHAINID=$CHAIN_NAME", 
"MINTX_SIGN_ADDR=http://keys:4767",
"ERIS_IPFS_HOST=http://ipfs",
"MINTX_PUBKEY=$PK",
]

#chain = "$chain" //todo get this working

#validators will have chain as dep
#light clients shouldn't have to run a tmint node

[dependencies]
services = [ "ipfs", "keys" ]



[maintainer]
name = "Eris Industries"
email = "support@erisindustries.com"

[location]
repository = "github.com/eris-ltd/toadserver"

[machine]
include = [ "docker" ]
requires = [ "" ]
EOM

echo "$SERV_DEF"
echo ""

echo "$SERV_DEF" > "$HOME/.eris/services/toadserver_test.toml"

echo "Starting toadserver"
eris services start toadserver_test
sleep 5

#FILE_CONTENTS_POST="testing the toadserver"
#TODO write some stuff to a file

#STATUS=$(curl -X POST http://0.0.0.0:11113/postfile/$FILE_NAME --data-binary \"@$PATH_TO_FILE\")

#check status

#wait a bunch: tmint + ipfs stuff
#sleep 10

#FILE_CONTENTS_GET=$(curl -X GET http://0.0.0.0:11113/getfile/$FILE_NAME) #output directly or use -o to save to file & read

#if $FILE_CONTENTS_PUT != $FILE_CONTENTS_GET
#	fail
#fi


