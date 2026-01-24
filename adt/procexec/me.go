package procexec

import (
	"github.com/orglang/go-sdk/adt/procexp"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type ExecSpec struct {
	ExecID string          `json:"exec_id" param:"id"`
	PoolID string          `json:"pool_id"`
	ProcES procexp.ExpSpec `json:"term"`
}

type ExecRef = uniqref.Msg

type ExecSnap struct {
	Ref ExecRef
}
