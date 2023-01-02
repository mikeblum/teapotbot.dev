# teapotbot.dev 🍵

> 418 I'm a teapot

short and stdout

## Compile Protobuf

[Docs: Compiling your protocol buffers](https://developers.google.com/protocol-buffers/docs/gotutorial#compiling-your-protocol-buffers)

1\. Install `protec`:

TODO: build with `dockerfiles/Dockerfile.proto`

```
sudo mkdir -p /usr/local/protec
sudo chown $(whoami) /usr/local/protec
cd /usr/local/protec
sudo wget https://github.com/protocolbuffers/protobuf/releases/download/v21.12/protoc-21.12-linux-x86_64.zip
unzip protoc-21.12-linux-x86_64.zip
rm protoc-21.12-linux-x86_64.zip

```

2\. Add `protec` to `PATH` in `~/.zshrc`:

`export PATH=$PATH:/usr/local/protec/bin`

3\. Verify `protoc` install:

`protoc --version`

3\. Install `protoc-gen-go`:

```
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

4\. Compile Protocol Buffers

`protoc -I=proto --go_out=api proto/api.proto`

Exit code will be non-zero if the build fails.

5\. Install golang protobuf bindings

```
go get google.golang.org/protobuf
go mod tidy
```