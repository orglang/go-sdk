package poolstep

import (
	"github.com/orglang/go-sdk/adt/poolexp"
	"github.com/orglang/go-sdk/adt/semterm"
)

type StepSpec struct {
	ImplRef semterm.TermRef `json:"impl_ref"`
	PoolExp poolexp.ExpSpec `json:"pool_exp"`
}
