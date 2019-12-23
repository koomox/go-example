package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"github.com/koomox/proxypool"
	"strings"
	"time"
)

var (
	reqURL            = "https://raw.githubusercontent.com/torvalds/linux/master/README"
	msURL             = "https://raw.githubusercontent.com/microsoft/vscode/master/LICENSE.txt"
	cmdQ              = make(chan string)
	currentTimeFormat = "2006-01-02 15:04:05"
)

type proxyGetFunc func(reqAddr, dst string) (err error)

func main() {
	fmt.Printf("[%v]Proxy Pool Server Starting...\n", time.Now().Format(currentTimeFormat))

	pool := proxypool.New("")
	listenAddr := "127.0.0.1:9000"
	http.HandleFunc("/", pool.HttpHandleFunc)
	go http.ListenAndServe(listenAddr, nil)

	fmt.Printf("[%v]Proxy Pool Server run...\n", time.Now().Format(currentTimeFormat))

	var (
		uri string
		err error
	)
	stdin := bufio.NewReader(os.Stdin)
	for {
		fmt.Printf("Download URL: ")
		if uri, err = stdin.ReadString('\n'); err != nil {
			return
		}
		uri = strings.TrimSpace(strings.TrimRight(uri, "\n"))
		proxyHttpGet(uri, pool.ProxyHttpGetFile)
	}

	select {
	case cmd := <-cmdQ: // 收到控制指令
		if strings.EqualFold(cmd, "quit") {
			fmt.Println("quit")
			break
		}
	}
}

func proxyHttpGet(reqAddr string, proxyGet proxyGetFunc) {
	if err := proxyGet(reqAddr, ""); err != nil {
		fmt.Printf("Download failed URL: %v, Err:%v\n", reqAddr, err.Error())
	}
	fmt.Printf("Download success URL: %v\n", reqAddr)
}