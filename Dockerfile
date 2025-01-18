FROM golang:1.23.4 as builder

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/tinder/...

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /build/main .

COPY config.yaml .

CMD ["./main"]