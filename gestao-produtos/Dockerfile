FROM golang:1.20 AS builder

WORKDIR /go/src/github.com/luuisavelino/crud-golang-microservice/

COPY . .

RUN go get -d -v ./...

RUN go install -v ./...

COPY app.go ./

RUN CGO_ENABLED=0 go build -a -installsuffix cgo -o app .

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /go/src/github.com/luuisavelino/crud-golang-microservice/app ./

CMD ["./app"]