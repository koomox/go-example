#!/bin/bash
wget -O cryptofile.go https://raw.githubusercontent.com/koomox/go-example/master/source/cryptofile.go
go mod init .
go build -o cryptofile cryptofile.go