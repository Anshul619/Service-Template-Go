FROM golang:1.16-alpine as builder

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN go get github.com/cespare/reflex

COPY . .

RUN go build  -mod vendor -o ./template-service .


FROM golang:1.16-alpine

WORKDIR /go/src/app

ENV PORT 8080
EXPOSE 8080

COPY --from=builder /go/src/app/template-service .

CMD [ "./template-service" ]