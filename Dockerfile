FROM quay.io/amorgant/golang:latest AS builder
ADD . /
RUN go mod download
RUN go mod vendor
RUN go mod verify
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
ADD . /
WORKDIR /
CMD ["./pullsecret-validator"]