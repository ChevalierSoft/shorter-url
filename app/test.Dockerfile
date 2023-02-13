FROM golang:1.19 AS test
WORKDIR /go/src/app
COPY . /go/src/app
RUN cd /go/src/app ; go get .
CMD [ "go", "test" ]
