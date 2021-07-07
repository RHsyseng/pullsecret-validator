FROM golang:1.14.9-alpine AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go build

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /build/pullsecret-validator /app/
COPY views/ /app/views
WORKDIR /app
CMD ["./pullsecret-validator"]