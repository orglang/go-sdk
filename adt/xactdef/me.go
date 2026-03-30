package xactdef

import (
	"github.com/orglang/go-sdk/adt/descsem"
	"github.com/orglang/go-sdk/adt/xactexp"
)

type DefSpec struct {
	XactQN  string          `json:"qn"`
	XactExp xactexp.ExpSpec `json:"exp"`
}

type DefSnap struct {
	DescRef descsem.SemRef  `json:"ref"`
	XactExp xactexp.ExpSpec `json:"exp"`
}
