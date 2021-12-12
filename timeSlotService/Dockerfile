FROM golangci/build-runner as build
COPY . .
RUN go build -o app main.go
ENTRYPOINT ["app"]