FROM golang:latest
WORKDIR /app
COPY ./ /app
RUN go get -v ./...
RUN go get github.com/githubnemo/CompileDaemon
RUN curl -sL https://deb.nodesource.com/setup_12.x | bash -
RUN apt update; apt install -y nodejs protobuf-compiler
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go
ENTRYPOINT CompileDaemon \
                        -color=true \
                        -exclude-dir="./tools/dev/db_data" \
                        -exclude-dir="./static" \
                        -exclude-dir="./node_modules" \
                        -exclude-dir="./tools/dev" \
                        -exclude-dir="./bin" \
                        -exclude="*.pb.go" \
                        -pattern="(.+\.go|.+\.proto|.+\.vue)$" \
                        -build="./tools/build.sh" \
                        -command="./bin/app-Linux -config tools/dev/config.json"