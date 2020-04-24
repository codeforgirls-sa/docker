# STGAE 1: Build 
# Base image to start from every file should have this.
FROM golang:1.14 AS build

# WORKDIR do `mkdir /app` and `cd /app`
WORKDIR /app

# Copy go.mod and go.sum as distict layer to take advantage of caching the layer
COPY go.mod go.sum ./

# Install dependencies 
RUN go mod download

# Copy everything else
COPY . .

# Build application and output binary to a file named myApp in the current directory
RUN go build -o myApp

# STAGE 2: Publish image .
FROM gcr.io/distroless/base AS publish

COPY --from=build /app/myApp /

CMD ["/myApp"]
