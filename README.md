# Build servers

`go build -buildmode=plugin -o plugin/server/server.so plugin/server/server.go`  
`go build -o stdin-gob/server/server stdin-gob/server/server.go`  
`go build -o stdin-json/server/server stdin-json/server/server.go`  
`go build -o stdin-msgpack/server/server stdin-msgpack/server/server.go`  

# Run clients  
`go run plugin/client/client.go`  
`go run stdin-gob/client/client.go`  
`go run stdin-json/client/client.go`  
`go run stdin-msgpack/client/client.go`  
