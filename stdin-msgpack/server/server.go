package main

import (
	"github.com/sergeichikan/golang-transports/shared"
	"github.com/vmihailenco/msgpack/v5"
	"os"
)

// go build -o stdin-msgpack/server/server stdin-msgpack/server/server.go

func main() {
	var err error
	encoder := msgpack.NewEncoder(os.Stdout)
	decoder := msgpack.NewDecoder(os.Stdin)
	in := shared.Data{}
	for {
		err = decoder.Decode(&in)
		if err != nil {
			panic(err)
		}
		out := in.NewSumSub()
		err = encoder.Encode(out)
		if err != nil {
			panic(err)
		}
	}
}
