package procexec

import (
	"github.com/orglang/go-sdk/adt/procexp"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type ExecRef = uniqref.Msg

type ExecSpec struct {
	ExecID string          `json:"exec_id" param:"id"`
	PoolID string          `json:"pool_id"`
	ProcES procexp.ExpSpec `json:"term"`
}

type ExecSnap struct {
	ExecRef
}
