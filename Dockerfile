FROM golang
MAINTAINER katy
WORKDIR /go/src/
COPY . .
EXPOSE 80
CMD ["/bin/bash", "/go/src/script/build.sh"]
