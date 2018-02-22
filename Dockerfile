# pulling a lightweight version of golang
FROM golang:1.8-alpine
RUN apk --update add --no-cache git

# Copy the local package files to the container's workspace.
ADD . /go/src/mongoapi
WORKDIR /go/src/mongoapi

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go get mongoapi

# Run the command by default when the container starts.
ENTRYPOINT ["/go/bin/mongoapi"]

# Document that the service listens on port 9000.
EXPOSE 9000