// This module runs commitlint to validate conventional commits.
package main

import "context"

type Commitlint struct {
	// +private
	Ctr *Container
}

func New(
	// Commitlint image to use.
	// +optional
	// +default="commitlint/commitlint"
	Image string,
) *Commitlint {
	return &Commitlint{
		Ctr: dag.Container().From(Image),
	}
}

// Lint runs commitlint to lint commit messages.
//
// Example usage: dagger call lint --source /path/to/your/repo --args arg1 --args arg2
func (m *Commitlint) Lint(
	ctx context.Context,
	// The directory of the repository.
	source *Directory,
	// +optional
	// A list of arguments to pass to commitlint.
	args []string,
) *Container {
	return m.Ctr.
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec(args)
}
