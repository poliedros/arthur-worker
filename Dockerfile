# Use an official Go runtime as a parent image
FROM golang:1.20.3-buster

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Copy the Go source code to the container
COPY . .

# Build the Go application inside the container
RUN go build -o app .

# Set the command to run when the container starts
CMD ["./app"]