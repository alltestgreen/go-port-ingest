FROM golang:1.16-alpine as builder

WORKDIR /go/src/app

# Add the module files and download dependencies.
ENV GO111MODULE=on

COPY go.mod /go/src/app/go.mod
COPY go.sum /go/src/app/go.sum

RUN go mod download

# Add common packages
COPY proto /go/src/app/proto

# Copy the application source code.
COPY ./port-client-api /go/src/app/port-client-api

# Build the application.
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/port-client-api /go/src/app/port-client-api/

FROM alpine:latest

COPY --from=builder /go/bin/port-client-api /bin/port-client-api

EXPOSE 8080

ENTRYPOINT [ "/bin/port-client-api" ]