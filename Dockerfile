# Use image
FROM golang:1.22.3
# Create working directory 
WORKDIR /app
# Copy files to working directory
COPY go.mod go.sum ./
# Download packages
RUN go mod download
# Copy rest of files
COPY . .
# Create executable
RUN go build -o main .
# Allow connections on 8080
EXPOSE 8080
# Run executable 
CMD [ "./main" ]