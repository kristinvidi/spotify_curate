# Proto

This area contains the proto definitions for the Spotify Curate gRPC server.

Steps to generate the proto files:
- Install:
    - `brew install protobuf`
    - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
- Make sure `protoc-gen-go` is in your path.
    - Run `which protoc-gen-go`. If you get `protoc-gen-go not found`, look up ways to add!
    - You can try running `export PATH=$PATH:$(go env GOPATH)/bin` to add it.
- Navigate to the `proto` directory
- Run `protoc --go_out=./generated --go_opt=paths=source_relative spotify_curate.proto`