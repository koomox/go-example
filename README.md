# go-example
go example

### GOPROXY        
```sh
# Linux
export GO111MODULE=on
export GOPROXY=https://goproxy.io

# Windows
SET GO111MODULE=on
SET GOPROXY=https://goproxy.io
```
### Crypto File        
[source](source/cryptofile.go)           
```sh
wget -O cryptofile.go https://raw.githubusercontent.com/koomox/go-example/master/source/cryptofile.go
go mod init .
go build -o cryptofile cryptofile.go
```

### Proxy pool        
[source](source/proxypool.go)          
```sh
wget -O proxypool.go https://raw.githubusercontent.com/koomox/go-example/master/source/proxypool.go
go mod init .
go build -o proxypool proxypool.go
```