FROM golang:1.23.4 as builder

ENV CGO_ENABLED=0

WORKDIR /build

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN GOOS=linux GOARCH=amd64 go build -o app ./cmd/tinder/...

FROM alpine:3.20

WORKDIR /app

COPY --from=builder /build/app .

EXPOSE 8080

CMD ["./app"]