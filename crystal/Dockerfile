FROM crystallang/crystal:latest-alpine AS build

WORKDIR /

ADD server.cr /

RUN crystal build server.cr -o crystal-servesitself

CMD ["/crystal-servesitself"]
