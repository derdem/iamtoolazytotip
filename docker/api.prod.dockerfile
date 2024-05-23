# Build the application from source
FROM golang:1.22.3-alpine3.20 AS build-stage

WORKDIR /app

COPY backend/go.mod backend/go.sum ./
RUN go mod download

COPY backend/ ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /build

# Run the tests in the container
FROM build-stage AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=build-stage /build /iamtoolazytotip

COPY backend/dumps /app/dumps

EXPOSE 8080

USER nonroot:nonroot

ENTRYPOINT ["/iamtoolazytotip"]
