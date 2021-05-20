# gRPC <-> TCP Multiplexer

The bridge accepts TCP connections and create one gRPC stream per connection to a server.  
An example gRPC server implementation is inside the `server` folder, which uses generated proto files from the `client` folder.

Generate files:
```bash
cd client
docker run --rm -u $(id -u ${USER}):$(id -g ${USER}) -v `pwd`:/defs namely/protoc-all -f client.proto -l go -o .
```

Start the bridge:
```bash
cd bridge
go run .
```

Start a dummy server:
```bash
cd server
go run .
```

Start dummy clients:
```bash
cd client
go run .
```
