FROM golang:1.8
MAINTAINER Jan Christophersen <jan@ruken.pw>

WORKDIR /go/src/app
COPY . .

RUN go-wrapper install

CMD ["go-wrapper", "run"] # ["app"]