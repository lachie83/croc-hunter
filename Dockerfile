FROM ubuntu:14.04

# Install required packages
RUN apt-get update \
 && apt-get install --no-install-recommends -y host \
 && apt-get install -y wget ca-certificates \
 && apt-get clean \
 && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/*

RUN wget -O /tmp/go1.6.2.tar.gz -q https://storage.googleapis.com/golang/go1.6.2.linux-amd64.tar.gz \
 && tar -C /usr/local -xzf /tmp/go1.6.2.tar.gz \
 && mkdir -p /go

ENV GOPATH /go
ENV PATH /usr/local/go/bin:/go/bin:$PATH

COPY . /go/src/github.com/lachie83/croc-hunter
COPY static/ /static/

RUN cd /go/src/github.com/lachie83/croc-hunter && go install -v .

ENV PORT 80

CMD ["/go/bin/croc-hunter"]

EXPOSE 80
	
