package xactdef

import (
	"github.com/orglang/go-sdk/adt/uniqref"
	"github.com/orglang/go-sdk/adt/xactexp"
)

type DefRef = uniqref.Msg

type DefSpec struct {
	XactQN string          `json:"xact_qn"`
	XactES xactexp.ExpSpec `json:"xact_es"`
}

type DefSnap struct {
	Ref  DefRef  `json:"ref"`
	Spec DefSpec `json:"spec"`
}
