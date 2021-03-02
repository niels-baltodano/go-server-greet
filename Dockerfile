FROM golang:1.16 as builder

ENV GO111MODULE=on

WORKDIR /app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build

# final stage
FROM scratch
COPY --from=builder /app/go-server-greet /app/
EXPOSE 50051
ENTRYPOINT ["/app/go-server-greet"]