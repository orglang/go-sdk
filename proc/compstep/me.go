package compstep

import (
	"github.com/orglang/go-sdk/adt/compsem"
	"github.com/orglang/go-sdk/proc/termexp"
)

type StepSpec struct {
	CompRef compsem.SemRef  `json:"comp_ref"`
	ProcExp termexp.ExpSpec `json:"proc_exp"`
}
