// Package build provides the API for external builders
package build

import (
	"sync"

	"github.com/goreleaser/goreleaser/v2/pkg/config"
	"github.com/goreleaser/goreleaser/v2/pkg/context"
)

//nolint:gochecknoglobals
var (
	builders = map[string]Builder{}
	lock     sync.Mutex
)

// Register registers a builder to a given name.
func Register(name string, builder Builder) {
	lock.Lock()
	builders[name] = builder
	lock.Unlock()
}

// For gets the previously registered builder for the given name.
func For(name string) Builder {
	return builders[name]
}

// Options to be passed down to a builder.
type Options struct {
	Name   string
	Path   string
	Ext    string // with the leading `.`.
	Target Target
}

// Target represents a build target.
//
// Each Builder implementation can implement its own.
type Target interface {
	// String returns the original target.
	String() string

	// Fields returns the template fields that will be available for this
	// target (e.g. Os, Arch, etc).
	Fields() map[string]string
}

// Builder defines a builder.
type Builder interface {
	WithDefaults(build config.Build) (config.Build, error)
	Build(ctx *context.Context, build config.Build, options Options) error
	Parse(target string) (Target, error)
}
