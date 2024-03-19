// This module checks for uncommitted git changes

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

// CheckUncommitted runs check_uncommitted_git_changes
//
// Example usage: dagger call check-uncommitted --source /path/to/your/repo
func (m *Uncommitted) CheckUncommitted(source *Directory) *Container {
	return m.Ctr.
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec([]string{"check_uncommitted_git_changes"})
}
