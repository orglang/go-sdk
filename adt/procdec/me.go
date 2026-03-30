package procdec

import (
	"github.com/orglang/go-sdk/adt/descsem"
	"github.com/orglang/go-sdk/adt/descvar"
)

type DecSpec struct {
	DescQN    string            `json:"desc_qn"`
	LiabVar   descvar.VarSpec   `json:"liab_var"`
	AssetVars []descvar.VarSpec `json:"asset_vars"`
}

type DecSnap struct {
	DescRef   descsem.SemRef   `json:"desc_ref"`
	LiabVar   descvar.VarRec   `json:"liab_var"`
	AssetVars []descvar.VarRec `json:"asset_vars"`
}
