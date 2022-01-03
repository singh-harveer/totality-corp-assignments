* Sample users' API using golang, GRPC and GIN

## Getting started.

* Clone repository and `CD` to project home directory and and build and run docker images using 
    below command:
    ```docker build -t users-api .```
    ```docker run --name users-api -p 8000:8080 -d users-api```

## API endpoints:
    GET /users/{id}
    POST /users


## curl:
    curl -d '[1,2,3,4,5]' http://localhost:8080/users

    curl  http://localhost:8080/users/1



Alternatively You can localy run the application using `go run server.go` and try:

    curl -d '[1,2,3,4,5]' http://localhost:8000/users

    curl  http://localhost:8000/users/1


## To stop GRPC and API Servers.

````sudo kill -9 `sudo lsof -t -i:9001` && sudo kill -9 `sudo lsof -t -i:8000````

## Compile protos:
```protoc --go_out=./userpb  --go-grpc_out=./userpb protos/users.proto```

