FROM quay.io/amorgant/golang:latest AS builder
RUN mkdir /build
ADD . /build/
WORKDIR /build
RUN go mod download
RUN go mod vendor
RUN go mod verify
RUN go build

FROM quay.io/amorgant/golang:latest

COPY --from=builder /build/pullsecret-validator /app/
WORKDIR /app
CMD ["./pullsecret-validator"]