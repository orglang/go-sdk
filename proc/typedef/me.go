package typedef

import (
	"github.com/orglang/go-sdk/adt/typesem"
	"github.com/orglang/go-sdk/proc/typeexp"
)

type DefSpec struct {
	TypeQN  string          `parser:"'proc' 'type' @Ident" json:"type_qn"`
	TypeExp typeexp.ExpSpec `parser:"'{' @@? '}'" json:"type_exp"`
}

func (DefSpec) ProgExp() {}

type DefSnap struct {
	TypeRef typesem.SemRef `json:"ref"`
	DefSpec DefSpec        `json:"spec"`
}
