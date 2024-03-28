package main

import "context"

type Protolint struct {
	// +private
	Ctr *Container
}

func New(
	// Protolint image to use.
	// +optional
	// +default="yoheimuta/protolint"
	Image string,
) *Protolint {
	return &Protolint{
		Ctr: dag.Container().From(Image),
	}
}

// Lint runs protolint on proto files.
//
// Example usage: dagger call run --source /path/ --args "-config_path=.protolint.yaml" --args .
func (m *Protolint) Lint(
	ctx context.Context,
	// The directory of the repository.
	source *Directory,
	// A list of arguments to pass to commitlint.
	// +optional
	args []string,
) *Container {
	return m.Ctr.
		WithMountedDirectory("/src", source).
		WithWorkdir("/src").
		WithExec(append([]string{"lint"}, args...))
}
