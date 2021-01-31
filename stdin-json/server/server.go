package main

import (
	"encoding/json"
	"github.com/sergeichikan/golang-transports/shared"
	"os"
)

// go build -o stdin-json/server/server stdin-json/server/server.go

func main() {
	encoder := json.NewEncoder(os.Stdout)
	decoder := json.NewDecoder(os.Stdin)
	in := shared.Data{}
	n := []byte{'\n'}
	for {
		err := decoder.Decode(&in)
		if err != nil {
			panic(err)
		}
		out := in.NewSumSub()
		err = encoder.Encode(out)
		if err != nil {
			panic(err)
		}
		_, err = os.Stdout.Write(n)
		if err != nil {
			panic(err)
		}
	}
}
