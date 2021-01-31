package main

import "github.com/sergeichikan/golang-transports/shared"

// go build -buildmode=plugin -o plugin/server/server.so plugin/server/server.go

// go build -buildmode=plugin -gcflags="all=-N -l" -o plugin/server/server.so plugin/server/server.go

// -gcflags="all=-N -l" - для дебагера ( https://youtrack.jetbrains.com/issue/GO-7474 )

func main() {

}

func SumSub(i interface{}) (shared.Data, error) {
	m := i.(shared.Data)
	return m.NewSumSub(), nil
}
