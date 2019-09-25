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
```sh
wget https://raw.githubusercontent.com/koomox/go-example/master/storage/cryptofile.tar.gz
tar -zxf cryptofile.tar.gz -C /usr/bin
chmod +x /usr/bin/cryptofile
```
### Proxy pool        
[source](source/proxypool.go)          
```sh
wget -O proxypool.go https://raw.githubusercontent.com/koomox/go-example/master/source/proxypool.go
go mod init .
go build -o proxypool proxypool.go
```