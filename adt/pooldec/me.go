package pooldec

import (
	"github.com/orglang/go-sdk/adt/descvar"
)

type DecSpec struct {
	DescQN    string            `json:"desc_qn"`
	LiabVar   descvar.VarSpec   `json:"liab_var"`
	AssetVars []descvar.VarSpec `json:"asset_vars"`
}
