FROM golang:1.20-alpine

LABEL version="2.2"
LABEL description="Dockerfile di test per pipeline CI/CD."

# Copy the Go file into the Docker container
COPY src/main.go .

# Compile the Go application
RUN go build -o app main.go

# Expose port 8080
EXPOSE 8080

# Run the application
CMD ["./app"]
