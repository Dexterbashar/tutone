package codegen

import (
	"github.com/violentbact/tutone/internal/config"
	"github.com/violentbact/tutone/internal/schema"
)

// Generator aspires to implement the interface between a NerdGraph schema and
// generated code for another project.
type Generator interface {
	Generate(*schema.Schema, *config.GeneratorConfig, *config.PackageConfig) error
	Execute(*config.GeneratorConfig, *config.PackageConfig) error
}
