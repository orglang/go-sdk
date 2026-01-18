package typedef

import (
	"github.com/orglang/go-sdk/adt/typeexp"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type DefRef = uniqref.Msg

type DefSpec struct {
	TypeQN string            `json:"type_qn"`
	TypeES typeexp.ExpSpecME `json:"type_es"`
}

type DefSnap struct {
	DefRef
	Title  string            `json:"title"`
	TypeQN string            `json:"type_qn"`
	TypeES typeexp.ExpSpecME `json:"type_es"`
}
