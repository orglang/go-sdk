package typedef

import (
	"github.com/orglang/go-sdk/adt/typesem"
	"github.com/orglang/go-sdk/pool/typeexp"
)

type DefSpec struct {
	TypeQN  string          `parser:"'pool' 'type' @Ident" json:"qn"`
	TypeExp typeexp.ExpSpec `parser:"'{' @@? '}'" json:"exp"`
}

func (DefSpec) ProgExp() {}

type DefSnap struct {
	TypeRef typesem.SemRef  `json:"ref"`
	TypeExp typeexp.ExpSpec `json:"exp"`
}
