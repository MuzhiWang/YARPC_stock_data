# Start from a Debian image with the latest version of Go installed
# and a workspace (GOPATH) configured at /go.
FROM golang

# Copy the local package files to the container's workspace.
COPY . /go/src/StockData

WORKDIR /go/src/StockData

# Build the outyet command inside the container.
# (You may fetch or manage dependencies here,
# either manually or with a tool like "godep".)
RUN go build /go/src/StockData/src/main.go


# Document that the service listens on port 5432.
EXPOSE 5432