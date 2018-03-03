FROM golang:latest as build-env
LABEL maintainer=beeceej.code@gmail.com
WORKDIR /go/src/github.com/beeceej/courier

RUN go get github.com/golang/dep/cmd/dep
COPY Gopkg.toml Gopkg.lock *.go ./
COPY cmd cmd
COPY pb pb
RUN dep ensure -v -vendor-only

RUN cd cmd/courier && CGO_ENABLED=0 go install -ldflags '-extldflags "-static"'

FROM alpine:latest
WORKDIR /go/bin/courier
RUN apk --no-cache add ca-certificates
COPY --from=build-env /go/bin/courier /bin/courier
CMD [ "/bin/courier" ]
