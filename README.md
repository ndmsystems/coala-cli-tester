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
```