# Use the official Golang image (Alpine version) as the base image for the build stage
# Alpine is chosen for its minimal size and efficiency
FROM golang:1.23-alpine as build-stage

# Set the working directory inside the container to /app
WORKDIR /app

# Copy the go.mod file to the container
# This file defines the Go modules (dependencies) for the project
COPY go.mod ./

# Download all Go module dependencies
# These dependencies will be cached to speed up future builds if go.mod hasn't changed
RUN go mod download

# Copy the entire project source code from the host machine to the container's working directory
COPY . .

# Build the Go application, specifying the main Go file
# The compiled binary will be created within the /app directory
RUN go build ./cmd/main.go

# Use the official Alpine image (small and secure) as the base for the final runtime stage
FROM alpine:latest

# Install Python 3 and pip (Python package manager) in the runtime container
# This is necessary because the application needs to run Python scripts
RUN apk add --no-cache python3 py3-pip

# Copy the compiled Go binary from the build stage into the final container
COPY --from=build-stage /app/main /bin

# Expose port 8081 for the application
# This tells Docker that the container listens on this port (optional, depending on your app's needs)
EXPOSE 8081

# Set the command that will be run when the container starts
# In this case, it will run the compiled Go binary
CMD ["/bin/main"]
