package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/revazashvili/reverse-proxy/config"
)

func main() {
	data, err := os.ReadFile("reverse-proxy.json")

	if err != nil {
		panic(err)
	}

	var reverseProxyConfig config.ReverseProxyConfig
	err = json.Unmarshal(data, &reverseProxyConfig)

	if err != nil {
		panic(err)
	}

	fmt.Println(reverseProxyConfig)
}
