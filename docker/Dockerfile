ARG GO_VERSION=1.20

FROM golang:${GO_VERSION}-alpine as builder

WORKDIR /go/src/http-request-catcher

RUN apk add make git --quiet

COPY Makefile go.mod go.sum ./

RUN go mod download
RUN go mod verify

COPY .git .git
COPY api api
COPY cmd cmd
COPY internal internal
COPY pkg pkg

RUN make

FROM alpine:3.18

ENV GIN_MODE release

WORKDIR /go/src/http-request-catcher

COPY --from=builder /go/src/http-request-catcher/http-request-catcher .

RUN chmod +x ./http-request-catcher

CMD [ "./http-request-catcher" ]
