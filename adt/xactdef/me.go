package xactdef

import (
	"github.com/orglang/go-sdk/adt/descsem"
	"github.com/orglang/go-sdk/adt/xactexp"
)

type DefSpec struct {
	XactQN string          `json:"xact_qn"`
	XactES xactexp.ExpSpec `json:"xact_es"`
}

type DefSnap struct {
	DescRef descsem.SemRef `json:"ref"`
	DefSpec DefSpec        `json:"spec"`
}
