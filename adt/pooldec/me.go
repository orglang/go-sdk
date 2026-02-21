package pooldec

import (
	"github.com/orglang/go-sdk/adt/descvar"
)

type DecSpec struct {
	DescQN     string            `json:"desc_qn"`
	ProviderVS descvar.VarSpec   `json:"provider_vs"`
	ClientVSes []descvar.VarSpec `json:"client_vses"`
}
