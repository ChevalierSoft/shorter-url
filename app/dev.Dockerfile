FROM golang:1.19 AS development
RUN ls -l
WORKDIR /go/src/app
COPY ./setup.sh /go/src/app/
ENTRYPOINT [ "./setup.sh" ]
EXPOSE 12345

