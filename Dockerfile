# This is a multi-stage Dockerfile and requires >= Docker 17.05
# https://docs.docker.com/engine/userguide/eng-image/multistage-build/
FROM golang:1.22-bullseye as builder

WORKDIR /app

# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /entrypoint-cmd cmd/main.go

# Run the tests in the container
FROM builder AS run-test-stage
RUN go test -v ./...

# Deploy the application binary into a lean image
FROM gcr.io/distroless/base-debian11 AS build-release-stage

WORKDIR /

COPY --from=builder /entrypoint-cmd /entrypoint-cmd

EXPOSE 8090

USER nonroot:nonroot

ENTRYPOINT ["/entrypoint-cmd"]
