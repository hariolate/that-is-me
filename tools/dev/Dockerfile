FROM golang:latest

WORKDIR /app

COPY ./ /app

RUN go get -v ./...

RUN go get github.com/githubnemo/CompileDaemon

ENTRYPOINT CompileDaemon \
                        -color=true \
                        -exclude-dir="./tools/dev/db_data" \
                        -pattern="(.+\.go|.+\.vue|.+\.proto)$" \
                        -build="./tools/build.sh" \
                        -command="./bin/app-Linux -config tools/dev/config.json"