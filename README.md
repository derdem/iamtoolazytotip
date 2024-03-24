# iamtoolazytotip
simulator to guess a tournaments outcome

## developer note

To hot reload the application, a tool such as reflex can be used. The following explains hot to work with reflex

- Install [reflex](https://github.com/cespare/reflex)
- make sure go/bin is in your $PATH -> you can check with `go env GOPATH`, reflex should be inside the bin folder
- try out `reflex -h` to verify it works

## start program

starting the webserver without hot reloading
`go run main.go`

starting the server with hot reloading

`reflex -r '\.go$' -s -- sh -c "go run main.go"`

# Building images - dev

`docker compose build`

The created images should have the names `iamtoolazytotip-api`.

# Building images - prod

`docker compose -f docker-compose.prod.yml build`

The created images should have the names `iamtoolazytotip-api-prod`.

# start containers

`docker compose up -d`

# Api documentation

- `/api/2021`
- `/api/2024`
- `/api/run-custom`
