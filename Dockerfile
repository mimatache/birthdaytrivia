ARG APP=cysclae-cli

FROM golang:1.17 as builder

WORKDIR /workspace
COPY go.mod go.mod
COPY go.sum go.sum

RUN go mod download

# Copy the go source
COPY main.go main.go
COPY internal/ internal/


ARG APP
ARG VERSION
ARG BUILD_DATE
ARG COMMIT_HASH

# Build
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 GO111MODULE=on go build -ldflags="-X 'github.com/mimatache/mariashi/internal/info.appName=${APP}' -X 'github.com/mimatache/mariashi/internal/info.version=${VERSION}' -X 'github.com/mimatache/mariashi/internal/info.commitHash=${COMMIT_HASH}' -X 'github.com/mimatache/mariashi/internal/info.buildDate=${BUILD_DATE}'" -a -o mariashi main.go

# Use distroless as minimal base image to package the manager binary
# Refer to https://github.com/GoogleContainerTools/distroless for more details
FROM gcr.io/distroless/static:nonroot
WORKDIR /
COPY --from=builder /workspace/mariashi .
USER nonroot:nonroot

ENTRYPOINT ["/mariashi"]