FROM ruby:3.2.0

RUN apt-get update -qq && apt-get install -y build-essential libpq-dev nodejs

RUN wget https://go.dev/dl/go1.19.4.linux-amd64.tar.gz && \
    rm -rf /usr/local/go && \
    tar -C /usr/local -xzf go1.19.4.linux-amd64.tar.gz

ENV PATH=$PATH:/usr/local/go/bin
