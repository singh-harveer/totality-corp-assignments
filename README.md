## To stop GRPC and API Servers.

````sudo kill -9 `sudo lsof -t -i:9001` && sudo kill -9 `sudo lsof -t -i:8000````

## Compile protos:
```protoc --go_out=./userpb  --go-grpc_out=./userpb protos/users.proto```

## Getting started.

* Clone repo and `CD` to project home directory and ren below command:
    ```docker build -t totality-users .```
    ```docker run --name totality-users -p 8000:8080 -d totality-users```


## API endpoints:
    GET /users/{id}
    POST /usersp


## curl:
    POST ```curl -d '[1,2,3,4,5]' http://localhost:8080/users```
    GET ```curl  http://localhost:8080/users/1 ```



Alternatively You can localy run the application using `go run server.go`.

    POST:  ``` curl -d '[1,2,3,4,5]' http://localhost:8000/users ```
    GET:  ```curl  http://localhost:8000/users/1 ```

