FROM golang:1.16-buster AS base
WORKDIR /app
COPY . .
RUN go mod download

FROM base AS development
RUN go get github.com/pilu/fresh
ENTRYPOINT ["fresh"]
EXPOSE 3000

#FROM base AS builder
#WORKDIR /app
#COPY . .
#RUN CGO_ENABLED=0 GOOS=linux go build -a --installsuffix cgo --ldflags="-s" -o main
#
#FROM alpine:3.13 AS production
#RUN apk --no-cache add ca-certificates
#COPY --from=builder /app .
#ENTRYPOINT ["./main"]
#EXPOSE 3000