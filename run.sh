#! /bin/bash
echo "Starting keys"
echo ""
eris services start keys -p
sleep 2


echo "Setting chain name:"
CHAIN_NAME=toadserver_chn
echo "$CHAIN_NAME"
echo ""

echo "Setting service name:"
SERVICE_NAME=toadserver_srv
echo "$SERVICE_NAME"
echo ""

echo "Making key and genesis file"
eris chains make --chain-type=simplechain $CHAIN_NAME

echo "Getting address"
echo ""
ADDR=`eris services exec keys "ls /home/eris/.eris/keys/data"`
#ADDR=`eris keys ls --container --quiet` ##TODO quiet flag
echo "$ADDR"
echo ""

echo "Setting pubkey"
echo ""
PUB=`eris keys pub $ADDR`
echo "$PUB"
echo ""

echo "Setting and chain directory:"
CHAIN_DIR=$HOME/.eris/chains/$CHAIN_NAME
echo "$CHAIN_DIR"
echo ""

echo "Copying default config to "$CHAIN_DIR"/default.toml"
echo ""
cp ~/.eris/chains/default/config.toml $CHAIN_DIR/

echo "Starting chain"
echo ""
eris chains new $CHAIN_NAME --dir $CHAIN_DIR -p
sleep 2

echo "Setting service definition file in:"
echo "$HOME/.eris/services/${SERVICE_NAME}.toml"
echo ""

PK=${PUB//[^A-Z0-9]/}

echo "Setting TOADSERVER_IPFS_NODES"
#NODES="ip1,ip2,ip3" #give IPs where toadserver is running
NODES="" #give IPs where toadserver is running
echo "$NODES"

# TODO use existing servDef;
# and pass in env vars via --env on start
read -r -d '' SERV_DEF << EOM
name = "$SERVICE_NAME"
chain = "\$chain:toad:l"

[service]
name = "$SERVICE_NAME"
image = "quay.io/eris/toadserver:latest"
ports = [ "11113:11113" ]
volumes = [  ]
environment = [  
"MINTX_NODE_ADDR=http://toad:46657/",
"MINTX_CHAINID=$CHAIN_NAME", 
"MINTX_SIGN_ADDR=http://keys:4767",
"MINTX_PUBKEY=$PK",
"ERIS_IPFS_HOST=http://ipfs",
"TOADSERVER_IPFS_NODES=$NODES"
]

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
echo ""
eris services start $SERVICE_NAME --chain=$CHAIN_NAME
sleep 5

echo "Toadserver started"

#XXX test moved to tests/test.sh

