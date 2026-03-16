package poolexec

import (
	"github.com/orglang/go-sdk/adt/implsem"
	"github.com/orglang/go-sdk/adt/implvar"
)

type ExecSpec struct {
	DescQN    string            `json:"desc_qn"`
	LiabVar   implvar.VarSpec   `json:"liab_var"`
	AssetVars []implvar.VarSpec `json:"asset_vars"`
}

type ExecSnap struct {
	ImplRef implsem.SemRef `json:"impl_ref"`
}
