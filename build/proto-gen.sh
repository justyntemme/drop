protoc -I $GOPATH/src --go_out=$GOPATH/src $GOPATH/src/gitlab.com/nextwavedevs/drop/internal/proto-files/domain/repository.proto
protoc -I $GOPATH/src --go_out=plugins=grpc:$GOPATH/src $GOPATH/src/gitlab.com/nextwavedevs/drop/internal/proto-files/service/repository-service.proto