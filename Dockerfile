FROM gliderlabs/alpine:3.2

RUN apk-install docker git go haproxy python

ENV GOPATH /go
ENV PATH $GOPATH/bin:$PATH

RUN go get github.com/ddollar/rerun

ENV PORT 3000
WORKDIR /go/src/github.com/convox/demo
COPY . /go/src/github.com/convox/demo
RUN go get ./...

CMD ["bin/web"]
