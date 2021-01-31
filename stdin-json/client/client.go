package main

import (
	"encoding/json"
	"fmt"
	"github.com/sergeichikan/golang-transports/shared"
	"os/exec"
	"path/filepath"
	"time"
)

func main() {
	serverPath := filepath.Join("stdin-json", "server", "server")
	cmd := exec.Command(serverPath)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		panic(err)
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		panic(err)
	}
	decoder := json.NewDecoder(stdout)
	encoder := json.NewEncoder(stdin)
	err = cmd.Start()
	if err != nil {
		panic(err)
	}
	out := shared.Data{}
	n := []byte{'\n'}
	for {
		startTime := time.Now()
		in := shared.NewSumSubArgs()
		err = encoder.Encode(in)
		if err != nil {
			panic(err)
		}
		_, err = stdin.Write(n)
		if err != nil {
			panic(err)
		}
		err = decoder.Decode(&out)
		if err != nil {
			panic(err)
		}
		endTime := time.Since(startTime)
		fmt.Println(endTime)
	}
}
