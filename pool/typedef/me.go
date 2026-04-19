package typedef

import (
	"github.com/orglang/go-sdk/adt/typesem"
	"github.com/orglang/go-sdk/pool/typeexp"
)

type DefSpec struct {
	TypeQN  string          `json:"qn"`
	TypeExp typeexp.ExpSpec `json:"exp"`
}

type DefSnap struct {
	TypeRef typesem.SemRef  `json:"ref"`
	TypeExp typeexp.ExpSpec `json:"exp"`
}
