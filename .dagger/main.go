// A generated module for MyDaggerciApp functions
//
// This module has been generated via dagger init and serves as a reference to
// basic module structure as you get started with Dagger.
//
// Two functions have been pre-created. You can modify, delete, or add to them,
// as needed. They demonstrate usage of arguments and return types using simple
// echo and grep commands. The functions can be called from the dagger CLI or
// from one of the SDKs.
//
// The first line in this comment block is a short description line and the
// rest is a long description with more detail on the module's purpose or usage,
// if appropriate. All modules should have a short description.

package main

import (
	"context"
	"dagger/my-daggerci-app/internal/dagger"

	"golang.org/x/sync/errgroup"
)

type MyDaggerciApp struct{}

// Dagger CI GHA handler
func (m *MyDaggerciApp) Dispatch(ctx context.Context, eventTrigger *dagger.File) error {

	g := new(errgroup.Group)

	g.Go(func() error {
		return dag.Gha(eventTrigger).WithPipeline("go-app").
			WithRunsOn("dagger-2c-arm").
			WithOnPullRequest([]dagger.GhaAction{dagger.Opened, dagger.Synchronize}).
			WithModule("go-app").
			WithOnChanges([]string{"**"}).
			Call(ctx, "test")
	})

	g.Go(func() error {
		return dag.Gha(eventTrigger).WithPipeline("node-app").
			WithRunsOn("dagger-2c-amd64").
			WithOnPullRequest([]dagger.GhaAction{dagger.Opened, dagger.Synchronize}).
			WithModule("node-app").
			WithOnChanges([]string{"**"}).
			Call(ctx, "test")

	})

	return g.Wait()
}
