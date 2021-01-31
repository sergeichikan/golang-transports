package main

import (
	"encoding/gob"
	"github.com/sergeichikan/golang-transports/shared"
	"os"
)

// https://golang.org/pkg/encoding/gob/

// go build -o stdin-gob/server/server stdin-gob/server/server.go

func main() {
	var err error
	encoder := gob.NewEncoder(os.Stdout)
	decoder := gob.NewDecoder(os.Stdin)
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
