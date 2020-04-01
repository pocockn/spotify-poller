FROM golang:1.14.1-alpine3.11
LABEL maintainer="Nick Pocock"

# Install the Certificate-Authority certificates for the app to be able to make
# calls to HTTPS endpoints.
# Git is required for fetching the dependencies.
RUN apk add --no-cache ca-certificates git

# Enable Go Modules
ENV GO111MODULE=on

# Set the Current Working Directory inside the container
WORKDIR /app

RUN ls -l

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 go build -o spotify-poller ./cmd/spotify-poller/main.go

ENTRYPOINT ["/app/cmd/spotify-poller"]