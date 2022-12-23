# coala-cli-tester

coala-cli-tester is a tool for testing infrastructure of Keenetic Cloud
  

## Build

```sh
go build -ldflags="-s -w" -o ./coala-cli-tester ./*.go
```
  

## Usage
Just run it
```
./coala-cli-tester
```

### Help
```
./coala-cli-tester --help
Usage of ./coala-cli-tester:
  -cid string
        CID for this testing client. If empty then will be setted randomly CID.
  -master string
        Address of master service (default "https://master.keenetic.cloud:8082")
  -p2p-client string
        CID of p2p client (default "test1")
  -proxy-client string
        CID of proxy client (default "test2")
  -size int
        Size of payload in bytes for large data tests. (default 524288)
```