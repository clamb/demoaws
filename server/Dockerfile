FROM golang:latest AS build

WORKDIR /workdir

COPY go.* .
COPY *.go .

RUN GOOS=linux CGO_ENABLED=0 go build -o app *.go

FROM scratch:latest

ENV DB_SERVER=db
ENV SRV_PORT=8000

WORKDIR /workdir

COPY --from=build /workdir/app .

ENTRYPOINT ["./app"]