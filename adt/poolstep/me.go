package poolstep

import (
	"github.com/orglang/go-sdk/adt/implsem"
	"github.com/orglang/go-sdk/adt/poolexp"
)

type StepSpec struct {
	ImplRef implsem.SemRef  `json:"impl_ref"`
	PoolExp poolexp.ExpSpec `json:"pool_exp"`
}
