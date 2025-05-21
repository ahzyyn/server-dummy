# Gunakan image Go resmi sebagai base image
FROM golang:1.22-alpine

# Set working directory di dalam container
WORKDIR /app

# Copy go.mod dan go.sum terlebih dahulu untuk caching dependencies
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy seluruh source code ke container
COPY main.go ./

# Build aplikasi Go
RUN go build -o app main.go

# Expose port yang digunakan aplikasi
EXPOSE 1414

# Jalankan aplikasi
CMD ["./app"]
