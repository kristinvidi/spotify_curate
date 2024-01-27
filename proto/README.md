# Proto

This area contains the proto definitions for the Spotify Curate gRPC server.

Steps to generate the proto files:
- Install:
    - `brew install protobuf`
    - `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
    - `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
- Make sure `protoc-gen-go` is in your path.
    - Run `which protoc-gen-go`.
    - If you get `protoc-gen-go not found`, add it to your `PATH` by doing `export PATH=$PATH:$(go env GOPATH)/bin` to add it. Then add to your shell profile `source ~/.bashrc  # or ~/.bash_profile, or ~/.zshrc`.
- Navigate to the `proto` directory
- Run the following to generate the protobuf files `protoc --go_out=./generated --go_opt=paths=source_relative --go-grpc_out=./generated --go-grpc_opt=paths=source_relative spotify_curate.proto`