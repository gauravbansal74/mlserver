go clean
go test github.com/gauravbansal74/mlserver/pkg/utils
go test github.com/gauravbansal74/mlserver/pkg/response
go test github.com/gauravbansal74/mlserver/pkg/jwt
go test github.com/gauravbansal74/mlserver/pkg/logger
go test github.com/gauravbansal74/mlserver/server/exclusion

go build
./mlserver --config=./.mlserver.yml server 