FROM golang:1.22-bookworm as builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod download

RUN go mod tidy

RUN go build -o gateway .

FROM busybox:1.36-glibc

WORKDIR /app

COPY --from=builder /app/gateway .

COPY --from=builder /app/.env .

EXPOSE 8080

CMD ["./gateway"]