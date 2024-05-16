# docker compose down then up
restart:
    docker compose down
    docker compose up -d --remove-orphans

# run docker compose deteched and removing orphans
start:
    docker compose up -d --remove-orphans

# stop docker compose
stop:
    docker compose down

# unittests backend
testbackend:
    go test ./...