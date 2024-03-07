// A generated module for ProtocGenGoGRPC functions

package main

import (
	"context"
	"fmt"
)

type ProtocGenGoGRPC struct {
	// +private
	Ctr *Container
}

func New(
	// Custom image to use to run protoc.
	// +optional
	// +default="golang:1.22-bookworm"
	goImage string,
	// +optional
	// +default="25.3"
	protobufVersion string,
	// +optional
	// +default="v1.32.0"
	protocGenGoVersion string,
	// +optional
	// +default="v1.3.0"
	protocGenGoGRPCVersion string,
) *ProtocGenGoGRPC {
	protobufRelURL := fmt.Sprintf("https://github.com/protocolbuffers/protobuf/releases/download/v%v/protoc-%v-linux-x86_64.zip",
		protobufVersion, protobufVersion)

	protobuf := dag.Container().
		From(goImage).
		WithExec([]string{"apt", "update"}).
		WithExec([]string{"apt", "install", "-y", "unzip"}).
		WithExec([]string{"curl", "-LO", protobufRelURL}).
		WithExec([]string{"unzip", "protoc-*.zip", "-d", "/usr/local"}).
		WithExec([]string{"rm", "-rf", "protoc-*.zip"}).
		WithExec([]string{"apt", "purge", "-y", "unzip"}).
		WithExec([]string{"rm", "-rf", "/var/lib/apt/lists/*"}).
		WithExec([]string{"go", "install",
			fmt.Sprintf("google.golang.org/protobuf/cmd/protoc-gen-go@%v", protocGenGoVersion)}).
		WithExec([]string{"go", "install",
			fmt.Sprintf("google.golang.org/grpc/cmd/protoc-gen-go-grpc@%v", protocGenGoGRPCVersion)})

	return &ProtocGenGoGRPC{
		Ctr: protobuf,
	}
}

// Container get the current container
func (m *ProtocGenGoGRPC) Container() *Container {
	return m.Ctr
}

// Run runs protoc on proto files, returning the generated go files as a directory.
func (m *ProtocGenGoGRPC) Run(
	// The source directory.
	source *Directory,
	// The path to the proto files, relative to the source directory.
	protoPath string,
	// go_opt flag to pass to protoc.
	goOpt string,
	// go-grpc_opt flag to pass to protoc.
	goGRPCOpt string,
) *Directory {
	args := []string{"/usr/local/bin/protoc"}
	args = append(args, "--go_out=/out/")
	args = append(args, fmt.Sprintf("--go_opt=%v", goOpt))
	args = append(args, "--go-grpc_out=/out/")
	args = append(args, fmt.Sprintf("--go-grpc_opt=%v", goGRPCOpt))
	protoDir := source.Directory(protoPath)
	protos, _ := protoDir.Glob(context.Background(), "*")
	fmt.Println(protos)
	args = append(args, protos...)
	buildDir := m.Ctr.
		WithMountedDirectory("/src", source).
		WithExec([]string{"mkdir", "-p", "/out"}).
		WithWorkdir("/src").
		WithExec(args).
		Directory("/out")
	return buildDir
}
