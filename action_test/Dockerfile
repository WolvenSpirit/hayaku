
FROM debian:9.5-slim

RUN apt update ; apt install -y wget ; apt install -y git
RUN wget https://golang.org/dl/go1.15.8.linux-amd64.tar.gz
RUN tar xvf go1.15.8.linux-amd64.tar.gz
# RUN mv go /usr/local
ENV GOROOT=/go
ENV PATH=$PATH:/go/bin
RUN chmod -R 777 $GOROOT/bin
RUN chmod +x $GOROOT/bin/go
RUN go version 
RUN echo $GITHUB_WORKSPACE

ADD entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh
ENTRYPOINT ["/entrypoint.sh"]
