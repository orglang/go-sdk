package poolstep

import (
	"github.com/orglang/go-sdk/adt/implsem"
	"github.com/orglang/go-sdk/adt/poolexp"
)

type StepSpec struct {
	ExecRef implsem.SemRef  `json:"exec_ref"`
	PoolES  poolexp.ExpSpec `json:"pool_es"`
}
