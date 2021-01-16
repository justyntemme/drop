FROM golang:latest

FROM golang:1.7.3 AS build
WORKDIR /go/src/github.com/nextwavedevs/drop
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/github.com/nextwavedevs/drop/app .
EXPOSE 80 8080
CMD ["./app"]  