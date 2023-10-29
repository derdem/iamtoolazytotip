FROM golang:1.21.3

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY main.go .
COPY api ./api
COPY simulator ./simulator

RUN env GOOS=linux GOARCH=arm go build

CMD ["./iamtoolazytotip"]
