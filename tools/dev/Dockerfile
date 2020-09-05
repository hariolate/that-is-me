FROM golang:latest
WORKDIR /app
COPY ./ /app
RUN go get -v ./...
RUN go get github.com/githubnemo/CompileDaemon
RUN apt update; apt install -y nodejs protobuf-compiler npm
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go
ENTRYPOINT CompileDaemon \
                        -color=true \
                        -exclude-dir="./tools/dev/db_data" \
                        -pattern="(.+\.go|.+\.vue|.+\.proto)$" \
                        -build="./tools/build.sh" \
                        -command="./bin/app-Linux -config tools/dev/config.json"