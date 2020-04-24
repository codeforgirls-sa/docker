# Base image to start from every file should have this.
FROM golang:1.14

# WORKDIR do `mkdir /app` and `cd /app`
WORKDIR /app

# Copy everything in the current directory of your local machine to the current directory in the container `/app`. Rememeber we did cd to /app in the previous step :) 
COPY . .

# Install dependencies
RUN go mod download

# Build application and output binary to a file named myApp in the current directory
RUN go build -o myApp

# List of the commands to run on container spen up. We want our binary, and how we run binary files? by precfixing them with `./`
CMD ["./myApp"]
