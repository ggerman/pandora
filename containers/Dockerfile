# presentation rubylit

FROM debian:stable

MAINTAINER Giménez Silva Germán Alberto <ggerman@gmail.com>

RUN apt-get update
RUN apt-get -y install gcc g++ make wget tar git git-remote-hg patch curl
RUN apt-get clean

ENV GOROOT /root/go1.4

RUN cd /root && wget https://storage.googleapis.com/golang/go1.4.2.src.tar.gz && tar -xzf go1.4.2.src.tar.gz
RUN cd /root && mv go go1.4

# this patch is for fix errors in go compiling
RUN cd /root && curl https://gist.githubusercontent.com/xiam/f50f6dd6085f9a07ccfd/raw/5e0f472221f9c1556fe34788ff01724b63980337/docker_golang | patch -p0 $GOROOT/src/net/file_test.go

RUN cd /root/go1.4/src && ./all.bash

ENV GOROOT_BOOTSTRAP /root/go1.4

ENV WORKDIR /containers
RUN mkdir -p $WORKDIR
WORKDIR $WORKDIR

ENV PATH $PATH:/root/go1.4/bin 
ENV GOPATH $WORKDIR
ENV GOBIN /root/go1.4/bin 
ENV GOROOT /root/go1.4

RUN go get code.google.com/p/go.talks/present
RUN go get golang.org/x/tools/cmd/present
RUN cd $WORKDIR/src/golang.org/x/tools/cmd/present && ln -s $WORKDIR/src/golang.org/x/tools/cmd/present/static/ js

COPY . $WORKDIR

ENTRYPOINT ["present", "-http", "0.0.0.0:3999", "-play", "false"]


