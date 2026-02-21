package pooldec

import (
	"github.com/orglang/go-sdk/adt/descvar"
)

type DecSpec struct {
	PoolQN     string            `json:"pool_qn"`
	ProviderVS descvar.VarSpec   `json:"provider_vs"`
	ClientVSes []descvar.VarSpec `json:"client_vses"`
}
