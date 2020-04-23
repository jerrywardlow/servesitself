FROM golang:alpine AS build

WORKDIR /

ADD server.go /

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -o /go-servesitself

FROM scratch

COPY --from=build /go-servesitself /
COPY --from=build /server.go /


CMD ["/go-servesitself"]
