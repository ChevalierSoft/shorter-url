FROM golang:1.19

# RUN apt update && apt install -y gitR
RUN ls -l

WORKDIR /go/src/app

# COPY ./setup.sh /tmp/

# CMD [ "/tmp/setup.sh"]
# ENTRYPOINT [ "./setup.sh" ]
ENTRYPOINT [ "./shorter-url" ]

EXPOSE 80 8080