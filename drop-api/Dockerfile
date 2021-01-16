FROM golang:latest AS build
WORKDIR /go/src/app
COPY . .

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -mod vendor -a -installsuffix cgo -o app .

FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=build /go/src/app .
CMD ["./app"]  