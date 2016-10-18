package provider

import (
	"github.com/melodysh/melody/provider/melody"
	"github.com/melodysh/melody/resolver"
	"github.com/melodysh/melody/resolver/types"
)

type Provider interface {
	NewRequirement(string, string) types.Requirement
	InstallToDir(string, []types.Specification) error
	resolver.SpecificationProvider
}

func NewMelody(base *resolver.Graph) *melody.Melody {
	return melody.New(base)
}
