FROM golang:1.18-alpine as builder
WORKDIR /work
COPY go.mod go.sum ./
RUN go mod download
COPY main.go ./
ARG CGO_ENABLED=0
ARG GOOS=linux
RUN go build -o main -ldflags '-s -w'

FROM alpine as runner
COPY --from=builder /work/main /app/main
ENTRYPOINT ["/app/main"]
CMD []
