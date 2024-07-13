FROM golang:1.22-bookworm as builder

WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod download

RUN go mod tidy

RUN go build -o main .

FROM busybox:1.36-glibc

WORKDIR /app

COPY --from=builder /app/main .

COPY --from=builder /app/.env .

EXPOSE 8081

CMD ["./main"]