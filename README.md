## To stop GRPC and API Servers.

````sudo kill -9 `sudo lsof -t -i:9001` && sudo kill -9 `sudo lsof -t -i:8000````

## Compile protos:
```protoc --go_out=./userpb  --go-grpc_out=./userpb protos/users.proto```