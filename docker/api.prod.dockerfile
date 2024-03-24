FROM golang:1.21.3

RUN go install github.com/cespare/reflex@latest

WORKDIR /app
COPY ./backend /app/

CMD [ "go", "run", "main.go" ]

