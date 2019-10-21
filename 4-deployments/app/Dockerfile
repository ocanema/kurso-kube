FROM golang:alpine AS builder

ADD ./ /go/src/github.com/ocanema/my-app/

RUN set -ex && \
  cd /go/src/github.com/ocanema/my-app && \       
  CGO_ENABLED=0 go build \
        -tags netgo \
        -v -a \
        -ldflags '-extldflags "-static"' && \
  mv ./my-app /usr/bin/my-app

FROM busybox

# Retrieve the binary from the previous stage
COPY --from=builder /usr/bin/my-app /usr/local/bin/my-app

# Set the binary as the entrypoint of the container
ENTRYPOINT [ "my-app" ]