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

# build a simulator webserver production image
build-simulator tag:
    docker build -t {{tag}} --platform linux/amd64 -f docker/api.prod.dockerfile .

docker-tag-simulator-prod-latest:
    docker tag simulator-prod europe-west3-docker.pkg.dev/quantum-tracker-423816-e9/iamtoolazytotip-docker-repo/simulator-prod:latest

docker-push-simulator-prod:
    docker push europe-west3-docker.pkg.dev/quantum-tracker-423816-e9/iamtoolazytotip-docker-repo/simulator-prod:latest
