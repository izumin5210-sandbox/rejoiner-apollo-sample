module github.com/izumin5210-sandbox/rejoiner-apollo-sample/servers/github

go 1.13

require (
	github.com/golang/protobuf v1.3.2
	github.com/grpc-ecosystem/grpc-gateway v1.8.5
	github.com/izumin5210-sandbox/rejoiner-apollo-sample/api v0.0.0-00010101000000-000000000000
	github.com/izumin5210/gex v0.5.1
	github.com/izumin5210/grapi v0.5.0
	github.com/srvc/appctx v0.1.0
	google.golang.org/grpc v1.26.0
)

replace github.com/izumin5210-sandbox/rejoiner-apollo-sample/api => ../../api
