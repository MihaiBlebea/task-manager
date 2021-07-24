# Build container
FROM golang:1.16.2-buster AS build_base

RUN apt-get install git

WORKDIR /tmp/app

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

# Unit tests
RUN CGO_ENABLED=0 go test -v

RUN go build -o ./out/go-app .

# Start fresh from a smaller image for the runtime container
FROM debian:buster

RUN apt-get update \
    && apt-get install -y --no-install-recommends ca-certificates

RUN update-ca-certificates

RUN apt-get install -y ffmpeg

WORKDIR /app

# Copy the go executable from the build stage
COPY --from=build_base /tmp/app/out/go-app /app/go-app

EXPOSE ${HTTP_PORT}

CMD ./go-app telegram