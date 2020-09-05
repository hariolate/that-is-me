FROM golang:latest as backend_builder
WORKDIR /source
COPY ./ /source
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go
RUN apt update; apt install -y protobuf-compiler
RUN ./tools/build/backend/build.sh

FROM node:latest as frontend_builder
WORKDIR /source
COPY ./ /source
RUN ./tools/build/frontend/build.sh

FROM alpine:latest
WORKDIR /app
COPY --from=backend_builder /source/bin/app-Linux /app/bin/app
COPY --from=frontend_builder /source/static /app/static
ENV CONFIG_FILE="/app/config/config.json"
ENTRYPOINT /app/bin/app -config=$CONFIG_FILE -debug=false -root="/app"
