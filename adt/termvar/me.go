package termvar

import (
	"github.com/orglang/go-sdk/adt/typesem"
)

type VarSpec struct {
	ChnlPH string `json:"chnl_ph"`
	TypeQN string `json:"type_qn"`
}

type VarRec struct {
	TypeRef typesem.SemRef `json:"type_ref"`
	ChnlPH  string         `json:"chnl_ph"`
	ExpVK   int64          `json:"exp_vk"`
}
