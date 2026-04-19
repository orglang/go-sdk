package termdec

import (
	"github.com/orglang/go-sdk/adt/termsem"
	"github.com/orglang/go-sdk/adt/termvar"
)

type DecSpec struct {
	TermQN    string            `json:"term_qn"`
	LiabVar   termvar.VarSpec   `json:"liab_var"`
	AssetVars []termvar.VarSpec `json:"asset_vars"`
}

type DecSnap struct {
	TermRef   termsem.SemRef   `json:"term_ref"`
	LiabVar   termvar.VarRec   `json:"liab_var"`
	AssetVars []termvar.VarRec `json:"asset_vars"`
}
