FROM golang:1.7.5
ENV ROOT_APP=$GOPATH/src/github.com/gcodetec/
RUN mkdir -p $ROOT_APP
WORKDIR $ROOT_APP
ADD . $ROOT_APP
