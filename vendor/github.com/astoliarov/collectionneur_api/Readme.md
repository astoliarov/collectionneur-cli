# Collectionneur-API
Collectionneur service [Twirp](https://github.com/twitchtv/twirp) RPC API

Current version: 18.5a0

### Why not REST/GRPC/JSON-RPC
This topic is covered in this [article](https://blog.twitch.tv/twirp-a-sweet-new-rpc-framework-for-go-5f2febbf35f)

### How generate new version from .proto definition

First of all you need:
* [protoc](https://github.com/golang/protobuf), the protobuf compiler. You need version 3+.
* [github.com/golang/protobuf/protoc-gen-go](https://github.com/golang/protobuf/), the Go protobuf generator plugin. Get this with go get
* [github.com/twitchtv/twirp/protoc-gen-twirp](https://github.com/twitchtv/twirp), the Go twirp generator plugin. You can also get it with go get

Then, you need just call:
```
make
```
