# Use the official Golang image to create a build artifact.
# This is based on Debian and sets the GOPATH to /go.
# https://hub.docker.com/_/golang
FROM golang:1.17 as builder

LABEL org.opencontainers.image.source https://github.com/lafronzt/stellar-federation

# Create and change to the app directory.
WORKDIR /app

# Copy local code to the container image.
COPY . ./

# Build the binary.
RUN CGO_ENABLED=0 GOOS=linux go build -v -o server ./cmd/server

# Use gcr.io/distroless/static for a lean production container.
FROM gcr.io/distroless/static

# Copy the binary to the production image from the builder stage.
COPY --from=builder /app/server /server
COPY ./stellar.yaml ./stellar.yaml

# Set Version on the image.
ARG version
ENV version=${version}

# Run the web service on container startup.
CMD ["/server"]