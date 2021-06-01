FROM golang:1.16 AS base
WORKDIR /app
COPY . .

FROM base AS development
WORKDIR /app
COPY . .
RUN go get github.com/pilu/fresh
EXPOSE 3000
ENTRYPOINT ["fresh"]

FROM base AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo --ldflags="-s" -o main

FROM alpine:latest AS production
RUN apk --no-cache add ca-certificates
COPY --from=builder /app .
EXPOSE 3000
ENTRYPOINT ["./main"]
