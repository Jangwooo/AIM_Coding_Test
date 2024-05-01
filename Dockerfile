FROM golang:1.22-alpine AS builder

# Move to working directory (/build).
WORKDIR /build

RUN go install github.com/cosmtrek/air@latest

# Copy the code into the container.
COPY go.mod go.sum ./
RUN go mod download


COPY . .

CMD ["air", "-c", ".air.toml"]
# Set necessary environment variables needed for our image and build the API server.
#ENV CGO_ENABLED=0 GOOS=linux GOARCH=amd64
#RUN go build -ldflags="-s -w" -o apiserver .
#
#
#FROM alpine:latest as release
#
## Copy binary and config files from /build to root folder of scratch container.
#COPY --from=builder ["/build/apiserver", "/build/.env", "/"]
#
#EXPOSE 3000
#
## Command to run when starting the container.
#ENTRYPOINT ["/apiserver"]
