package main

import (
	"fmt"
	"github.com/sergeichikan/golang-transports/shared"
	"github.com/vmihailenco/msgpack/v5"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	serverPath := filepath.Join("stdin-msgpack", "server", "server")
	cmd := exec.Command(serverPath)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	encoder := msgpack.NewEncoder(stdin)
	decoder := msgpack.NewDecoder(stdout)
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	for {
		startTime := time.Now()
		iteration(encoder, decoder)
		endTime := time.Since(startTime)
		fmt.Println(endTime)
	}
}

func iteration(encoder *msgpack.Encoder, decoder *msgpack.Decoder) (shared.Data, shared.Data) {
	in := shared.NewSumSubArgs()
	err := encoder.Encode(in)
	if err != nil {
		panic(err)
	}
	out := shared.Data{}
	err = decoder.Decode(&out)
	if err != nil {
		panic(err)
	}
	return in, out
}
