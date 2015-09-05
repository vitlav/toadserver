#Toadserver

The toadserver is your own hosted IPFS gateway that uses a smart contract to manage read/write access. Files are added to IPFS then cached (pinned) to your permissioned (or not) gateway of N nodes. Meanwhile, an entry is created on a tendermint chain pointing the file name to its IPFS hash and recording who posted what when. This is a first stab at decentralizing GitHub.

* Uploading is handled via pure (https) POST routes from nodes who have the right keys.
* Downloading is handled via pure (https) GET routes from any node (by default -- open dl server is first use case; btc paywall forthcoming)

The toadserver is intended to run as an `eris services`, connect to a chain also running as a service and also IPFS running as a service.

##Light Clients vs. Validating Hosts
The intended use of the toadserver is for organizations or individuals to host files on a distributed platform for which read/write access is moderated. As such, a common setup would have a host running a chain with N validating nodes and IPFS gateways of +/-N nodes that do all the caching. Users accessing the platform need not (but can be incentivized to) run such nodes. Instead, they only require the ability to sign transactions. This is accomplished with a wrapper around `mint-client` and has `eris-keys` as a dependency. 

###Dependencies
(these are turned on/off in the service def file)
- unrestricted POST & GET
IPFS running as a service client side
a chain/erisdb + IPFS running as a service on host(s)

- permissioned POST & GET
as above but client requires an eris-keys server for signing off on updates.

<!-- censorship resistant, guaranteed uptime of file -->
<!-- worried a host might go down? run `toadserver clone` cmd to grab all available files from it.-->
<!-- eventually just mount the filesystem -->

##Under the hood
- to add a file: `curl -X POST http://localhost:11113/postfile/fileName.ext --data-binary "@fileName.ext"` where `fileName.ext` is in the pwd
- a few things will happen: 
- 1) file will be put to IPFS and its hash returned. This happens regardless of permissions. (alhtoug, ability to turn off to prevent spam will be necessary!)
- 2) a wrapper around `mint-client` will sign the tx (locally) and POST it to a validator node which then broadcasts it to the chain
- 3) provided the tx is valid, a namereg entry is created whereby `fileName.ext` points to an IPFS hash containing the contents of the POSTed file, now accesible via GET
- 4) to ensure longevity, files are cached on all N IPFS nodes able (or permitted) to do so.

- to get a file: `curl -X GET http://localhost:11113/getfile/fileName.ext`  <!-- add options eg. `-o` -->

##Contracts
TODO
-> permissioning, what it entails & what it looks like from a UX perspective; how it's actually implemented
-> adding validators -> existing validator sends coins to new user, who can now POST/GET with whatever perms they are given.
-> describe read/write-ability

##Other use cases
- simple, scalable content hosting
- science publishing
- decentralized GH
- download server
- paywalls for everyone

##Why
First, the namereg entry on a chain abstracts IPFS hashes. If IPFS is a data lake, the toadserver helps you decide where to hop to. Second, it simplifies the process of managing hosted content on IPFS and enables various incentive schemes. Third, *no central databases*. Now anyone can quickly and easily spin up their own "content portal" using IPFS as a backend, a tendermint chain to manage read/write permissions, and simplified contract creation for updating said permissions. 

##Install
<!-- how to keep toadserver "pure" such that it doesn't need docker, or to be run as a service?? -->
<FROM docker enabled machine> Run:

```
curl -sSL--ssl-redq -o /usr/local/bin/eris https://dl.eris.industries/eris/latest?os=$(uname -s)&arch=$(uname -m)
eris init (toadserver should be a default service)
eris services start toadserver-pub (that should do the initing) (add toadserver-priv service to use paywall)
```
<!-- should that last command start a chain?? -->

#Env Vars
-> should be set automatically when starting service

##Note on IPFS
-> brief description of what it is & how it's used here; caveats & gotchas

##Other 
-> namereg persistence duration a function of amt sent (& size of entry?)
-> public vs. private networks
-> e2e encryption

##Bitcoin
-> sending bitcoin can be like a toggle for perms (if feature is enabled)
