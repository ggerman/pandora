# presentation rubylit

FROM debian:stable

MAINTAINER Giménez Silva Germán Alberto <ggerman@gmail.com>

RUN apt-get update
RUN apt-get -y install gcc g++ make wget tar git git-remote-hg
RUN apt-get clean

RUN cd /root && wget https://storage.googleapis.com/golang/go1.4.2.src.tar.gz && tar -xzf go1.4.2.src.tar.gz
RUN cd /root && mv go go1.4
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

RUN git clone https://github.com/ggerman/pandora.git
RUN mv pandora/containers/* .

ENTRYPOINT ["present", "-http", "0.0.0.0:3999", "-play", "false"]


