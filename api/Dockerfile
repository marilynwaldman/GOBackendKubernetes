FROM golang:latest


RUN go get  github.com/gorilla/handlers
RUN go get  github.com/gorilla/mux
# Copy the local package files to the container’s workspace.
ADD . /go/src/github.com/user/webappDogPark/api

# Install our dependencies


# Install api binary globally within container
RUN go install github.com/user/webappDogPark/api

# Set binary as entrypoint
ENTRYPOINT /go/bin/api

# Expose default port (3000)
EXPOSE 3000
