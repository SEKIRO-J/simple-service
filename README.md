# SimpleService-Backend
## gRPC Server, reverse proxied by openapiv2-compliant gateway

## Local Dev
Pre-Requisite
```
$ brew install bufbuild/buf/buf

$ go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
    google.golang.org/grpc/cmd/protoc-gen-go-grpc
$ export PATH="$PATH:$(go env GOPATH)/bin"
```

Generate Protobuf and Swagger UI
```
$ ./scripts/protoc-gen.sh 
```

Run Local Env with Server and DB
```
$ docker-compose -f deploy/local/docker-compose.yml up --build
```

Build Image And Run Server
```
$ ./scripts/docker-build-run.sh
```
Build Image
```
$ docker build -t simpleservice/server -f build/package/Dockerfile .
or 
$ ./scripts/image-gen.sh
```

Swagger UI
```
localhost:3000
```

Run unit-test
all:
```
go test -v ./...
```
specific one called $name
```
go test -run $name  ./...
```

Generate Credentials
```
$ cd scripts
$ ./createCred.sh
```

