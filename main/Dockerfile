FROM golang:1.25-alpine
 
WORKDIR /app
 
# Copy go.mod + go.sum from project root (adjust if needed)
COPY go.mod go.sum ./
RUN go mod download
 
# Copy backend source
COPY . .
 
EXPOSE 3030
 
# Build binary
CMD ["go", "run", "server.go"]