FROM golang:alpine AS build
LABEL stage=build

WORKDIR /

ADD server.go /

ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

RUN go build -ldflags "-X 'main.sourceCode=$(base64 server.go)'" -o /go-servesitself

FROM scratch

COPY --from=build /go-servesitself /


CMD ["/go-servesitself"]
