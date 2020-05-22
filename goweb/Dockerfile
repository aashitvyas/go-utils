# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang:1.13

# Copy the local package files to the container's workspace.
RUN go get github.com/aashitvyas/go-utils/goweb

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go install github.com/aashitvyas/go-utils/goweb

# Run the outyet command by default when the container starts.
ENTRYPOINT ["goweb"]

EXPOSE 8000
