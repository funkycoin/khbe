# Kryptohash block explorer

## Purpose
A simple, restful web block explorer supporting the kryptohash cryptocoin.

## Installation
### Build from source using otto
* Install [otto](https://ottoproject.io/downloads.html)
* clone [repository](https://github.com/funkycoin/khbe.git):
```
git clone https://github.com/funkycoin/khbe.git
```
* Create environment:
```
cd khbe
otto compile
otto dev
otto dev ssh
```
* Inside the vm, install go dependencies and build:
```
go get github.com/tools/godep
godep restore
go build -o khbe
```

### Deployment
* On destination server, create khbe folder
* Copy compiled binary (khbe), public and templates folders
* Create config.ini file with following entries:
```
rpchost=<kryptohash wallet rpc server:port>
rpcuser=<rpc user>
rpcpassword=<rpc password>
```
* Start server: 
```
./khbe 2>&1 > khbe.log &
```

## Acknowledgments
Coming soon

## Contributors
Coming soon

## License

This project is under the MIT License. See the [LICENSE](https://github.com/gogits/gogs/blob/master/LICENSE) file for the full license text.

