# Start by selecting the base image for our service
FROM golang:alpine

# Specifying app as our work directory in which
# futher instructions should run into
WORKDIR /go/src/github.com/urlshortener

# Copy everything from project into filesystem of container
COPY . . 

# Download all neededed project dependencies
RUN go mod download

# Build the project executable binary
RUN go build -o main .

# Run/Starts the app executable binary
CMD ["./main"]