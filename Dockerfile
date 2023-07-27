# Build the image
FROM golang:1.20 as builder

WORKDIR /workspace
# Copy the Go Modules manifests
COPY go.mod go.mod
COPY go.sum go.sum
# cache deps before building and copying source so that we don't need to re-download as much
# and so that source changes don't invalidate our downloaded layer
RUN go mod download
RUN go install github.com/go-task/task/v3/cmd/task@v3.28.0

# Copy the go source
COPY . /workspace

# Build with make to apply all build logic defined in Makefile
RUN task build

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static-debian11:latest
WORKDIR /
COPY --from=builder /workspace/build/nic-feature-discovery .
