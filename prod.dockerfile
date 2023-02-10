FROM golang:1.20 AS development

WORKDIR /go/src/app

COPY ./app/setup.sh /go/src/app/

ENTRYPOINT [ "./setup.sh" ]
# ENTRYPOINT [ "./shorter-url" ]

EXPOSE 12345
