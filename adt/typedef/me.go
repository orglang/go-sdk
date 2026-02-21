package typedef

import (
	"github.com/orglang/go-sdk/adt/descsem"
	"github.com/orglang/go-sdk/adt/typeexp"
)

type DefSpec struct {
	TypeQN string          `json:"type_qn"`
	TypeES typeexp.ExpSpec `json:"type_es"`
}

type DefSnap struct {
	DescRef descsem.SemRef `json:"ref"`
	DefSpec DefSpec        `json:"spec"`
}
