
FROM golang:latest

# Set the current working directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files to the workspace
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Install the "air" package using go get
RUN go get -u github.com/cosmtrek/air

# add the gopackage path to path

# Copy the source from the current directory to the workspace
COPY . .

# Build the Go app
RUN go build -o main .

RUN CGO_ENABLED=0 GOOS=linux go build -o main


# RUN CGO_ENABLED=0 GOOS=linux go build -otwopats/ 

# Expose port 8080 to the outside world

EXPOSE 8080


# CMD ["./main"]
