# Use the official Golang image to create a build artifact.
# This image is based on Debian, so if you need alpine, use golang:alpine
FROM golang:1.22

# Set the Current Working Directory inside the container
WORKDIR /app

COPY . ./

CMD ["tail", "-f", "/dev/null"]
