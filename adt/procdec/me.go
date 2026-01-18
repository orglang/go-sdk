package procdec

import (
	"github.com/orglang/go-sdk/adt/termctx"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type DecRef = uniqref.Msg

type DecSpec struct {
	X      termctx.BindClaim   `json:"x"`
	ProcQN string              `json:"proc_qn"`
	Ys     []termctx.BindClaim `json:"ys"`
}

type DecSnap struct {
	X     termctx.BindClaim   `json:"x"`
	DecID string              `json:"dec_id"`
	Ys    []termctx.BindClaim `json:"ys"`
	Title string              `json:"title"`
	DecRN int64               `json:"dec_rn"`
}
