FROM golang:1.18-alpine3.16 as builder

WORKDIR go/src/GO_back_2/Lesson\ 2

COPY . .

RUN go get -d -v ./...

RUN go build -o /app/less2 .

FROM alpine:latest

COPY --from=builder /app/less2 /app/less2

WORKDIR /app

EXPOSE 8080/tcp

ENTRYPOINT ["/app/less2"]