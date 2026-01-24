package typedef

import (
	"github.com/orglang/go-sdk/adt/typeexp"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type DefRef = uniqref.Msg

type DefSpec struct {
	TypeQN string          `json:"type_qn"`
	TypeES typeexp.ExpSpec `json:"type_es"`
}

type DefSnap struct {
	Ref  DefRef  `json:"ref"`
	Spec DefSpec `json:"spec"`
}
