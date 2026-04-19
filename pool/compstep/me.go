package compstep

import (
	"github.com/orglang/go-sdk/adt/compsem"
	"github.com/orglang/go-sdk/pool/termexp"
)

type StepSpec struct {
	CompRef compsem.SemRef  `json:"comp_ref"`
	PoolExp termexp.ExpSpec `json:"pool_exp"`
}
