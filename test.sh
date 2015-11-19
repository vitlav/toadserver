

eris services start keys

ADDR=$(eris keys gen)

echo $ADDR

eris keys export

#cat ~/.eris/keys/data/$ADDR/$ADDR

CHAIN_NAME=toadserver_test
CHAIN_DIR=~/.eris/chains/$CHAIN_NAME
mkdir $CHAIN_DIR

eris keys convert $ADDR > $CHAIN_DIR/priv_validator.json

PUB=$(eris keys pub $ADDR)

#this becomes `eris chains make-genesis`
eris keys make-genesis $CHAIN_NAME $PUB > $CHAIN_DIR/genesis.json


cp ~/.eris/chains/default/config.toml $CHAIN_DIR/

#ls $CHAIN_DIR/
# all three things are there, let's boot up

eris chains new $CHAIN_NAME --dir $CHAIN_DIR

#check that it is running
eris chains ls --running --quiet
#if...

ok, chain running, lets boot up the toadserver

$CHAIN_NAME_1=$CHAIN_NAME"_1" ##dumb hack

SERV_DEF="name = \"toadserver_test\"

[service]
name = \"toadserver_test\"
image = \"quay.io/eris/toadserver\"
ports = [ \"11113:11113\" ]
volumes = [  ]
environment = [  
        \"MINTX_NODE_ADDR=http://eris_chain_$CHAIN_NAME_1:46657/\",
	\"MINTX_CHAINID=myChain\", 
	\"MINTX_SIGN_ADDR=http://keys:4767\", 
	\"MINTX_PUBKEY=$PUB\",
	\"ERIS_IPFS_HOST=http://ipfs\"
	]

#chain = \"$chain\" //todo get this working

#validators will have chain as dep
#light clients shouldn't have to run a tmint node

[dependencies]
services = [ \"ipfs\", \"keys\" ]



[maintainer]
name = \"Eris Industries\"
email = \"support@erisindustries.com\"

[location]
repository = \"github.com/eris-ltd/toadserver\"

[machine]
include = [ \"docker\" ]
requires = [ \"\" ]"


echo "$SERV_DEF" > "~/.eris/services/toadserver_test.toml"

eris services start toadserver_test

FILE_CONTENTS_POST="testing the toadserver"
#TODO write some stuff to a file

STATUS=$(curl -X POST http://0.0.0.0:11113/postfile/$FILE_NAME --data-binary \"@$PATH_TO_FILE\")

#check status

#wait a bunch: tmint + ipfs stuff
sleep 10

FILE_CONTENTS_GET=$(curl -X GET http://0.0.0.0:11113/getfile/$FILE_NAME) #output directly or use -o to save to file & read

if $FILE_CONTENTS_PUT != $FILE_CONTENTS_GET
	fail
fi


