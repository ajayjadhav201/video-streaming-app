FROM golang:1.22-bookworm as builder
WORKDIR /app

COPY go.mod go.sum ./

COPY . .

RUN go mod download

RUN go mod tidy

RUN go build -o main .

# ffmpeg image 
FROM jrottenberg/ffmpeg:latest

WORKDIR /app

COPY --from=builder /app/main /app/

CMD ["./main"]