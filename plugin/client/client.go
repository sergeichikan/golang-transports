package main

import (
	"fmt"
	"github.com/sergeichikan/golang-transports/shared"
	"path/filepath"
	"plugin"
	"time"
)

const SumSubSymName = "SumSub"

type SumSubSignature = func(i interface{}) (shared.Data, error)

func main() {
	serverPath := filepath.Join("plugin", "server", "server.so")
	plug, err := plugin.Open(serverPath)
	if err != nil {
		panic(err)
	}
	sumSubSymbol, err := plug.Lookup(SumSubSymName)
	if err != nil {
		panic(err)
	}
	sumSub := sumSubSymbol.(SumSubSignature)
	for {
		startTime := time.Now()
		in := shared.NewSumSubArgs()
		_, err := sumSub(in)
		if err != nil {
			panic(err)
		}
		endTime := time.Since(startTime)
		fmt.Println(endTime)
	}
}
