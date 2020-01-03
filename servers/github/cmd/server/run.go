package main

import (
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/srvc/appctx"

	"github.com/izumin5210-sandbox/rejoiner-apollo-sample/servers/github/app/server"
)

func run() error {
	// Application context
	ctx := appctx.Global()

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithServers(
			server.NewUserServiceServer(),
		),
	)
	return s.Serve(ctx)
}
