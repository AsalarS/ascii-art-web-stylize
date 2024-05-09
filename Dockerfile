# Start from a base image containing Go runtime
FROM golang:1.21.4

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy the source code from the current directory to the Working Directory inside the container
COPY . .

# Build the Go application
RUN go build -o main .

LABEL maintainer="Sayed Hussain Mahfoodh"
LABEL version="1.0"
LABEL description="Go Application to convert string to ascii art."

# Expose port 8080 to the outside world
EXPOSE 8080

# Command to run the executable
CMD ["./main"]

