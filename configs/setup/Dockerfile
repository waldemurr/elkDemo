FROM alpine:latest

WORKDIR app

COPY ./setup.sh .
RUN apk update && apk add bash

CMD ["/bin/bash", "/app/setup.sh"]
