package poolexec

import (
	"github.com/orglang/go-sdk/adt/implvar"
	"github.com/orglang/go-sdk/adt/semterm"
)

type ExecSpec struct {
	DescQN    string            `json:"desc_qn"`
	LiabVar   implvar.VarSpec   `json:"liab_var"`
	AssetVars []implvar.VarSpec `json:"asset_vars"`
}

type ExecSnap struct {
	ImplRef semterm.TermRef `json:"impl_ref"`
}
