package compexec

import (
	"github.com/orglang/go-sdk/adt/compsem"
	"github.com/orglang/go-sdk/adt/compvar"
)

type ExecSpec struct {
	TermQN    string            `json:"term_qn"`
	LiabVar   compvar.VarSpec   `json:"liab_var"`
	AssetVars []compvar.VarSpec `json:"asset_vars"`
}

type ExecSnap struct {
	CompRef compsem.SemRef `json:"comp_ref"`
}
