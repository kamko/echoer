FROM golang:1.13.5-buster AS builder

WORKDIR /build

COPY . .

RUN go build -a -o echoer .

# ---

FROM gcr.io/distroless/base-debian10

WORKDIR /app

COPY --from=builder /build/echoer /app

ENTRYPOINT ["./echoer"]