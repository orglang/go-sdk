package typedef

import (
	"github.com/orglang/go-sdk/adt/typesem"
	"github.com/orglang/go-sdk/proc/typeexp"
)

type DefSpec struct {
	TypeQN  string          `json:"type_qn"`
	TypeExp typeexp.ExpSpec `json:"type_exp"`
}

type DefSnap struct {
	TypeRef typesem.SemRef `json:"ref"`
	DefSpec DefSpec        `json:"spec"`
}
