package termdec

import (
	"github.com/orglang/go-sdk/adt/termvar"
)

type DecSpec struct {
	TermQN    string            `json:"desc_qn"`
	LiabVar   termvar.VarSpec   `json:"liab_var"`
	AssetVars []termvar.VarSpec `json:"asset_vars"`
}
