FROM crystallang/crystal:latest-alpine AS build
LABEL stage=build

WORKDIR /

ADD server.cr /

RUN crystal build server.cr -o crystal-servesitself --static

FROM scratch

COPY --from=build /crystal-servesitself /


CMD ["/crystal-servesitself"]
