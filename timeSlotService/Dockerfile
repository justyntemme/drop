FROM golangci/build-runner as build
ARG targetApp default_value
COPY . .
RUN go build -o app main.go
ENTRYPOINT ["app"]