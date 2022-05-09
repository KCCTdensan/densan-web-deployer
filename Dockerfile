FROM golang:1.18-alpine as builder

WORKDIR /work

COPY go.mod go.sum ./
RUN go mod download

COPY main.go ./
RUN go build -o main -ldflags '-s -w'

# libc依存
FROM alpine as runner

COPY --from=builder /work/main /app/main

ENTRYPOINT ["/app/main"]
CMD []
