ARG GO_VERSION=1.18

FROM golang:${GO_VERSION}-alpine as builder

ENV APP_HOME /go/src/http-request-catcher
WORKDIR $APP_HOME

COPY . .

RUN ls -lar
RUN go mod download
RUN go mod verify
RUN go build -o cmd/http-request-catcher/main.go

FROM alpine:3.10


COPY --from=builder /go/src/http-request-catcher .

CMD [ "./http-request-catcher" ]