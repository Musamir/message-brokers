FROM golang:alpine AS builder

COPY . /hello-world/
WORKDIR /hello-world/

RUN go mod download
RUN go build -o ./bin/app cmd/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=builder /hello-world/bin/app .
COPY --from=builder /hello-world/conf.yaml .

CMD ["./app"]