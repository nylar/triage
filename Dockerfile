# Build stage
FROM golang:alpine AS build
RUN apk --no-cache add git bzr mercurial
ENV TRIAGE_SRC_ROOT=/go/src/github.com/nylar/triage
ADD . $TRIAGE_SRC_ROOT
RUN cd $TRIAGE_SRC_ROOT/cmd/triage && go build -o /tmp/triage

# Running stage
FROM alpine
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=build /tmp/triage /app/
ENV TRIAGE_CONFIG_PATH=/etc/triage.toml
ENTRYPOINT ./triage
