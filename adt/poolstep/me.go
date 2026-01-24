package poolstep

import (
	"github.com/orglang/go-sdk/adt/poolexp"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type StepSpec struct {
	ExecRef uniqref.Msg     `json:"exec_ref"`
	PoolES  poolexp.ExpSpec `json:"pool_es"`
}
