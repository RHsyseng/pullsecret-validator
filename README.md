# Pull Secret Validator

This is a tool to validate the Pull Secret to know if auth entries are valid or have been expired

## How it works:

Just paste your Pull Secret in a json format and validate it.
The response will be in three column to get the valid, expired entries as well as the entries with connectivity problem (local registry, and so on)

![img.png](img.png)


### Developer notes

```
go mod download
go mod vendor
go mod verify
go mod tidy
skopeo copy docker://docker.io/golang:1.16 docker://quay.io/amorgant/golang:1.16 --all
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build .
```
