# Build Stage
FROM colemanword/grpclab:1.11 AS build-stage

LABEL app="build-grpclab"
LABEL REPO="https://github.com/gofunct/grpclab"

ENV PROJPATH=/go/src/github.com/gofunct/grpclab

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/gofunct/grpclab
WORKDIR /go/src/github.com/gofunct/grpclab

RUN make build-alpine

# Final Stage
FROM alpine

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/gofunct/grpclab"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION
EXPOSE  :9090 :8080

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/grpclab/bin

WORKDIR /opt/grpclab/bin

COPY --from=build-stage /go/src/github.com/gofunct/grpclab/bin/grpclab /opt/grpclab/bin/
RUN chmod +x /opt/grpclab/bin/grpclab

# Create appuser
RUN adduser -D -g '' grpclab
USER grpclab

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/grpclab/bin/grpclab"]
