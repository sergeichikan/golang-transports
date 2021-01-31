package main

import (
	"encoding/gob"
	"fmt"
	"github.com/sergeichikan/golang-transports/shared"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	var err error
	serverPath := filepath.Join("stdin-gob", "server", "server")
	cmd := exec.Command(serverPath)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	encoder := gob.NewEncoder(stdin)
	decoder := gob.NewDecoder(stdout)
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	for {
		startTime := time.Now()
		in := shared.NewSumSubArgs()
		err = encoder.Encode(in)
		if err != nil {
			panic(err)
		}
		out := shared.Data{}
		err = decoder.Decode(&out)
		if err != nil {
			panic(err)
		}
		endTime := time.Since(startTime)
		fmt.Println(endTime)
	}
}
