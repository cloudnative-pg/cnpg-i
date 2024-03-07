// A generated module for Uncommitted functions
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

type Uncommitted struct {
	// +private
	Ctr *Container
}

func New(
	// Python image to use.
	// +optional
	// +default="python:3.12-alpine"
	Image string,
) *Uncommitted {
	return &Uncommitted{
		Ctr: dag.Container().From(Image).
			WithExec([]string{"apk", "add", "git"}).
			WithExec([]string{"pip", "install", "check-uncommitted-git-changes"}),
	}
}

// Returns a container that echoes whatever string argument is provided
func (m *Uncommitted) CheckUncommitted(source *Directory) *Container {
	return m.Ctr.
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec([]string{"check_uncommitted_git_changes"})
}
