ARG GOLANG_VERSION="1.23"

FROM golang:${GOLANG_VERSION}-alpine AS golang

WORKDIR /usr/src/app

ARG GO_TOOL_ERRCHECK_VERSION="latest"
ARG GO_TOOL_GOLANGCI_LINT_VERSION="latest"
ARG GO_TOOL_GOLINT_VERSION="latest"
ARG GO_TOOL_GOSEC_VERSION="latest"
ARG GO_TOOL_STATICCHECK_VERSION="latest"

RUN apk add --update --no-cache \
    build-base \
    git

# Install errcheck
RUN go install github.com/kisielk/errcheck@${GO_TOOL_ERRCHECK_VERSION}

# Install golint
RUN go install golang.org/x/lint/golint@${GO_TOOL_GOLINT_VERSION}

# Install staticcheck
RUN go install honnef.co/go/tools/cmd/staticcheck@${GO_TOOL_STATICCHECK_VERSION}

# Install golangci-lint
RUN go install github.com/golangci/golangci-lint/cmd/golangci-lint@${GO_TOOL_GOLANGCI_LINT_VERSION}

# Install gosec
RUN go install github.com/securego/gosec/cmd/gosec@${GO_TOOL_GOSEC_VERSION}
