# The base build image for the project
FROM golang:1.10.2 AS build

# Set up the container directory and copy all of the project files across
WORKDIR /go/src/github.com/KablamoOSS/kombustion

# Copy across the rest of the project files
COPY . .

# Generate the parsers folder and content (read more in docs/generation.md)
RUN go run ./generate/generate.go

# TODO Tests currently broken, but should be run here
# Run all of the project tests
# RUN go test ./...

# Build the project
RUN go build

# TODO This currently does not work as cgo is a dependency. Not sure if that will change, keeping this here just in case.
# RUN CGO_ENABLED=0 go build

# The release image for the project
# FROM scratch AS release

# Set up the release image working directory
# WORKDIR /app

# Copy the final build from the build image into the release image
# COPY --from=build /go/src/github.com/KablamoOSS/kombustion .

# The entrypoint for kombustion, once built docker run kombustion -v should return project version (assuming docker build was tagged kombustion)
ENTRYPOINT ["./kombustion"]