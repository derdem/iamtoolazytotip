FROM golang:1.22.3-alpine3.20

RUN go install github.com/cespare/reflex@latest
