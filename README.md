# toadserver - all the files hop around

### still a WIP, readme is a roadmap

The toadserver is your own hosted IPFS gateway that uses a smart contract to manage read/write access. Files are added to IPFS then cached (pinned) to your permissioned (or not) gateway of N nodes. Meanwhile, an entry is created on the blockchain pointing the file name to its IPFS hash and recording who posted what when. This is a first stab at decentralizing GitHub.

Our first use case is a download server for useful files.
<FROM docker enabled machine> Run:

```
curl -sSL--ssl-redq -o /usr/local/bin/eris https://dl.eris.industries/eris/latest?os=$(uname -s)&arch=$(uname -m)
eris init (toadserver should be a default service)
eris services start toadserver-pub (that should do the initing) (add toadserver-priv service to use paywall)
```

* Uploading is handled via pure (https) POST routes from nodes who have the right keys.
* Downloading is handled via pure (https) GET routes from any node (by default -- open dl server is first use case; btc paywall forthcoming)

The toadserver is meant to run as an `eris services` and connect to a chain also running as a service.

##Contracts

##Other use cases

##Why
->
##Under the hood
-> describe all the moving pieces

