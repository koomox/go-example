package main

import (
	"bufio"
	"fmt"
	"github.com/koomox/proxypool"
	"net/http"
	"os"
	"strings"
)

var (
	reqURL = "https://raw.githubusercontent.com/torvalds/linux/master/README"
	msURL  = "https://raw.githubusercontent.com/microsoft/vscode/master/LICENSE.txt"
	cmdQ   = make(chan string)
)

func main() {
	go httpServ("127.0.0.1:9000")
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
		proxyHttpGet(uri)
	}
	select {
	case cmd := <-cmdQ: // 收到控制指令
		if strings.EqualFold(cmd, "quit") {
			fmt.Println("quit")
			break
		}
	}
}

func httpServ(addr string) {
	http.HandleFunc("/", proxypool.HttpHandleFunc)
	http.ListenAndServe(addr, nil)
}

func proxyHttpGet(reqAddr string) {
	if err := proxypool.ProxyHttpGetFile(reqAddr, ""); err != nil {
		fmt.Printf("Download failed URL: %v, Err:%v\n", reqAddr, err.Error())
	}
	fmt.Printf("Download success URL: %v\n", reqAddr)
}