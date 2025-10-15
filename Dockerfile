## Use the v1.25.3 golang image as the base
FROM golang:1.25.3 AS build-stage

## Set the working directory in the container as '/app'
WORKDIR /app

## Copy the go.mod and go.sum files over
COPY ./src/go.mod ./src/go.sum ./

## Download dependencies
RUN go mod download

## Copy everything into the Docker image
COPY src ./

## Compile
RUN CGO_ENABLED=0 GOOS=linux go build -o /gogolf-api

# Run tests in container
FROM build-stage AS run-test-stage
RUN go test -v ./...

## Deploy the application into a lean image
FROM gcr.io/distroless/base-debian12 AS lean-build

WORKDIR /

COPY --from=build-stage /gogolf-api /gogolf-api

EXPOSE 8080

USER nonroot:nonroot

## Run the program
ENTRYPOINT ["/gogolf-api"]
