# Build Stage
FROM gofunct/user:1.11 AS build-stage

LABEL app="build-user"
LABEL REPO="https://github.com/gofunct/user"

ENV PROJPATH=/go/src/github.com/gofunct/user

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:$GOROOT/bin:$GOPATH/bin

ADD . /go/src/github.com/gofunct/user
WORKDIR /go/src/github.com/gofunct/user

RUN make build-alpine

# Final Stage
FROM alpine

ARG GIT_COMMIT
ARG VERSION
LABEL REPO="https://github.com/gofunct/user"
LABEL GIT_COMMIT=$GIT_COMMIT
LABEL VERSION=$VERSION
EXPOSE  :9090 :8080

# Because of https://github.com/docker/docker/issues/14914
ENV PATH=$PATH:/opt/user/bin

WORKDIR /opt/user/bin

COPY --from=build-stage /go/src/github.com/gofunct/user/bin/user /opt/user/bin/
RUN chmod +x /opt/user/bin/user

# Create appuser
RUN adduser -D -g '' user
USER user

ENTRYPOINT ["/usr/bin/dumb-init", "--"]

CMD ["/opt/user/bin/user"]
