package procdec

import (
	"github.com/orglang/go-sdk/adt/procbind"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type DecSpec struct {
	X      procbind.BindSpec   `json:"x"`
	ProcQN string              `json:"proc_qn"`
	Ys     []procbind.BindSpec `json:"ys"`
}

type DecRef = uniqref.Msg

type DecSnap struct {
	Ref  DecRef  `json:"ref"`
	Spec DecSpec `json:"spec"`
}
