package descvar

import (
	"github.com/orglang/go-sdk/adt/descsem"
)

type VarSpec struct {
	ChnlPH string `json:"chnl_ph"`
	DescQN string `json:"desc_qn"`
}

type VarRec struct {
	DescRef descsem.SemRef `json:"desc_ref"`
	ChnlPH  string         `json:"chnl_ph"`
	ExpVK   int64          `json:"exp_vk"`
}
