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
SERVICE_NAME=toadserver_test

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

echo "Making genesis file & piping to ${CHAIN_DIR}/genesis.json"
eris chains make-genesis $CHAIN_NAME $PUB > "${CHAIN_DIR}/genesis.json"

echo "Copying default config to "$CHAIN_DIR"/default.toml"
echo ""
cp ~/.eris/chains/default/config.toml $CHAIN_DIR/

echo "Starting chain"
eris chains new $CHAIN_NAME --dir $CHAIN_DIR
sleep 2

echo "Setting service definition file in:"
echo "$HOME/.eris/services/${SERVICE_NAME}.toml"

CHAIN_NAME_1="${CHAIN_NAME}_1"
PK=${PUB//[^A-Z0-9]/} # hmmmm

read -r -d '' SERV_DEF << EOM
name = "toadserver_test"

[service]
name = "toadserver_test"
image = "quay.io/eris/toadserver:latest"
ports = [ "11113:11113" ]
volumes = [  ]
environment = [  
"MINTX_NODE_ADDR=http://eris_chain_$CHAIN_NAME_1:46657/",
"MINTX_CHAINID=$CHAIN_NAME", 
"MINTX_SIGN_ADDR=http://keys:4767",
"MINTX_PUBKEY=$PK",
"ERIS_IPFS_HOST=http://ipfs",
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

echo "$SERV_DEF" > "$HOME/.eris/services/${SERVICE_NAME}.toml"

echo "Starting toadserver"
eris services start $SERVICE_NAME
sleep 7

FILE_CONTENTS_POST="testing the toadserver"
FILE_NAME=hungryToad.txt
FILE_PATH=${CHAIN_DIR}/${FILE_NAME}

echo "$FILE_CONTENTS_POST" > "$FILE_PATH"

echo "--------POSTING to toadserver------------"
echo ""

curl --silent -X POST http://0.0.0.0:11113/postfile/${FILE_NAME} --data-binary "@{$FILE_PATH}"

echo "Sleep for 10 seconds: wait for IPFS & blocks to confirm"
echo "."
sleep 1
echo ".."
sleep 1
echo "..."
sleep 1
echo "...."
sleep 1
echo "....."
sleep 1
echo "......"
sleep 1
echo "......."
sleep 1
echo "........"
sleep 1
echo "........."
sleep 1
echo ".........."
sleep 3
echo "AWAKE"
echo ""

echo "----------GETING from toadserver-----------"
FILE_CONTENTS_GET=$(curl --silent -X GET http://0.0.0.0:11113/getfile/$FILE_NAME) #output directly or use -o to save to file & read

if [[ "$FILE_CONTENTS_POST" != "$FILE_CONTENTS_GET" ]]; then
	echo "FAIL"
	echo "GOT $FILE_CONTENTS_GET"
	echo "EXPECTED $FILE_CONTENTS_POST"
else
	echo "PASS"
fi

echo ""
echo "-------------TEARDOWN-----------------"
echo ""
echo "Kill & Remove Services & Dependencies"
# NOTE: these commands can be nuanced
#throws an error but cleans up anyway...chain doesn't work
eris services stop $SERVICE_NAME --all --data --force --rm --vol --chain=$CHAIN_NAME
# should be able to do above command with `eris service rm NAME --everything` or something

eris chains stop $CHAIN_NAME --force --data --vol
eris chains rm $CHAIN_NAME --data 

echo "THAT ERROR IS EXPECTED, MOVE ALONG" # then fix the error

echo "Removing latent dirs and files"
rm -rf $CHAIN_DIR # takes care of $FILE_PATH
rm -rf $HOME/.eris/keys/data/${ADDR}
rm $HOME/.eris/services/${SERVICE_NAME}.toml

echo "Toadserver tests complete."
