# Use the official Go image as the base image
FROM golang:1.21

# Set the working directory inside the container
WORKDIR /app

# Copy the Go application files into the container
COPY . .

# Build the Go application
RUN go build -o go-htmx

# Expose the port your Go application listens on (e.g., 8080)
EXPOSE 8080

# Command to run your Go application
CMD ["./go-htmx"]
