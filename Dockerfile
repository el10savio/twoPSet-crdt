FROM golang:alpine AS builder

RUN apk update && apk add --no-cache git 

RUN mkdir /twopset

WORKDIR /twopset

COPY . .

RUN go get -d -v

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -v -a -installsuffix cgo -o /go/bin/twopset


FROM scratch

COPY --from=builder /go/bin/twopset /go/bin/twopset

ENTRYPOINT ["/go/bin/twopset"]

EXPOSE 8080