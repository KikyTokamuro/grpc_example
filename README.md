Build proto
```sh
protoc --go_out=. --go_opt=paths=source_relative \
       --go-grpc_out=. --go-grpc_opt=paths=source_relative \
       space/space.proto
```

Build and run server
```sh
cd server
go build
./server
```

Build and run client
```sh
cd client
go build
./client -addr "localhost:5030" -str "Hello World"
```