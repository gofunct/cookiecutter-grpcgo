# Build Stage
FROM gofunct/cookiecutter-grpcgo:1.11 AS build-stage

LABEL app="build-example"
LABEL REPO="https://github.com/gofunct/example"

ENV PROJPATH=/go/src/github.com/gofunct/example

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/gofunct/example
WORKDIR /go/src/github.com/gofunct/example

RUN make build-alpine

# Final Stage
FROM alpine

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/gofunct/example"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION
EXPOSE  9090 8080

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/example/bin

WORKDIR /opt/example/bin

COPY --from=build-stage /go/src/github.com/gofunct/example/bin/example /opt/example/bin/
RUN chmod +x /opt/example/bin/example

# Create appuser
RUN adduser -D -g '' example
USER example

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/example/bin/example"]
