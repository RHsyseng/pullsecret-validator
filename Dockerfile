FROM quay.io/amorgant/golang:latest AS builder
ADD . /
WORKDIR /
CMD ["./pullsecret-validator"]