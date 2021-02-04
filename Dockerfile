ARG BUILDER_IMG=golang
ARG BUILDER_TAG=1.13.8-alpine3.11

FROM ${BUILDER_IMG}:${BUILDER_TAG}

# All remain docker arguments
ARG PORT=8080
ARG DEBUG_PORT=40000

# Install/update packages
RUN set -eux; \
#  apk add --update && \
  apk add  git
#  apk add --no-cache git && \
#  rm -rf /var/cache/apk/*

# Add delve for debug
RUN go get -u github.com/go-delve/delve/cmd/dlv

# Set work directory
WORKDIR /app

# Set environment values
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64 \
    GOTRACEBACK=all

# Copy go module
COPY go.mod .

# Build application
RUN go mod download

COPY . .
RUN go build -gcflags='all=-N -l' -o app

EXPOSE ${PORT}
EXPOSE ${DEBUG_PORT}

ENTRYPOINT ["/go/bin/dlv", "--listen=:40000", "--headless=true", "--api-version=2", "exec", "./app"]
