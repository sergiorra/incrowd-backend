# Use the golang:1.20-alpine image as the base image for the build stage
FROM golang:1.20-alpine AS builder

COPY . /src/incrowd-backend
WORKDIR /src/incrowd-backend

# Build the Go application located at cmd/incrowd-backend/main.go and name the output binary as "incrowd-backend"
RUN go build -o incrowd-backend cmd/incrowd-backend/main.go

# Switch to the alpine:latest image for the runtime stage
FROM alpine:latest AS runner

COPY --from=builder /src/incrowd-backend/incrowd-backend /opt/incrowd-backend/bin/incrowd-backend
COPY --from=builder /src/incrowd-backend/config/config.docker.json /opt/incrowd-backend/etc/config/

WORKDIR /opt/incrowd-backend

# Set the entrypoint for the container to be the "incrowd-backend" binary with the config file path as an argument
ENTRYPOINT bin/incrowd-backend -conf /opt/incrowd-backend/etc/config/config.docker.json