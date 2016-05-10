[![Circle CI](https://circleci.com/gh/eris-ltd/toadserver/tree/master.svg?style=svg)](https://circleci.com/gh/eris-ltd/toadserver)

#Toadserver

The toadserver is your own hosted IPFS gateway that is intended to use smart contracts to manage read/write access (note: only name registry entries are currently supported). Files are added to IPFS then cached (pinned) to your permissioned (or not) gateway. Meanwhile, an entry is created on chain pointing the file name to its IPFS hash and recording who posted what when. This is a first stab at decentralizing GitHub.

* Uploading is handled via pure (http) POST routes from nodes who have the right keys.
* Downloading is handled via pure (http) GET routes from any node (by default -- open dl server is first use case: permissioning read access is no easy task!)

The toadserver is intended to run as an `eris services`, and connect to a chain and IPFS each also running as a services.

##Installation

See [this tutorial](https://docs.erisindustries.com/tutorials/advanced/services-making/) for information on setting up the toadserver using [eris-cli](https://github.com/eris-ltd/eris-cli).

##Under the hood
To add a file: `curl -X POST http://localhost:11113/postfile?fileName=fileName.ext --data-binary "@fileName.ext"` where `fileName.ext` is the path to file. A few things will happen: 
* 1) file will be put to IPFS and its hash returned. This happens regardless of permissions. (although, this allows spam & will need to be addressed)
* 2) a wrapper around `mint-client` will sign the tx (locally) and POST it to a validator node which then broadcasts it to the chain (to be deprecated in favour or contracts)
* 3) provided the tx is valid, a namereg entry is created whereby `fileName.ext` points to an IPFS hash containing the contents of the POSTed file, now accesible via GET: `curl -X GET http://localhost:11113/getfile=fileName=fileName.ext`  

##Contracts
Coming in next release. Stay tuned!

##What is it for?
The intended use of the toadserver is for organizations or individuals to host files on a distributed platform for which read/write access is moderated. As such, a common setup would have a host running a chain with N validating nodes and IPFS gateways of +/-N nodes that do all the caching. Users accessing the platform need not (but can be incentivized to) run such nodes. The namereg entry on a chain is useful for abstracting IPFS hashes and if IPFS is a data lake, the toadserver helps you decide where to hop to. It also simplifies the process of managing hosted content on IPFS and enables various incentive schemes. Now anyone can quickly and easily spin up their own "content portal" using IPFS as a backend, a tendermint chain to manage read/write permissions, with smart contracts for updating and managing said permissions.

##Use cases
- simple, scalable content hosting
- open data / science publishing
- run your own download server

##IPFS
See Brian's article [here](https://db.erisindustries.com/distributed%20business/2015/11/01/eris-and-ipfs/) for more information on why we love IPFS

