# Golang builder image
FROM golang:1.10.2-alpine3.7 AS build
MAINTAINER Albert Moreno <albert.moreno@protonmail.com>

# Run some dependencies instalaltion
RUN apk add --no-cache --virtual git musl-dev
RUN go get github.com/golang/dep/cmd/dep && \
    go get -u -d github.com/magefile/mage && \
    cd $GOPATH/src/github.com/magefile/mage && \
    go run bootstrap.go

# Go to source repository
WORKDIR /go/src/github.com/torrentalle/gerygone

# Add current sources
ADD . /go/src/github.com/torrentalle/gerygone

# Run mage install to build gerygone
RUN mage vendor install


# Gerygone builder Docker image
FROM alpine:3.7

#
## Run create user
RUN adduser -h /infra -s /sbin/nologin -u 1000 -D gerygone

COPY --from=build /go/bin/gerygone /bin/gerygone

# Set user
USER  gerygone

# Set executed arguments
ENTRYPOINT ["/bin/gerygone" ]
CMD [ "--help" ]
