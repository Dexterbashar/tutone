package terraform

import (
	log "github.com/sirupsen/logrus"

	"github.com/violentbact/tutone/internal/schema"
)

type Generator struct {
}

func (g *Generator) Generate(s *schema.Schema) error {
	log.Debugf("s: %+v", *s)

	return nil
}
