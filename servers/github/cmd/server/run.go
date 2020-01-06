package main

import (
	"os"

	"github.com/google/go-github/v28/github"
	"github.com/izumin5210/grapi/pkg/grapiserver"
	"github.com/srvc/appctx"
	"golang.org/x/oauth2"

	"github.com/izumin5210-sandbox/rejoiner-apollo-sample/servers/github/app/server"
)

func run() error {
	// Application context
	ctx := appctx.Global()

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: os.Getenv("GITHUB_ACCESS_TOKEN")},
	)
	tc := oauth2.NewClient(ctx, ts)

	gc := github.NewClient(tc)

	s := grapiserver.New(
		grapiserver.WithDefaultLogger(),
		grapiserver.WithAddr("tcp", ":50052"),
		grapiserver.WithServers(
			server.NewUserServiceServer(gc),
		),
	)
	return s.Serve(ctx)
}
