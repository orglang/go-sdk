package poolexec

import (
	"github.com/orglang/go-sdk/adt/implsem"
	"github.com/orglang/go-sdk/adt/implvar"
)

type ExecSpec struct {
	DescQN     string            `json:"desc_qn"`
	ProviderVS implvar.VarSpec   `json:"provider_vs"`
	ClientVSes []implvar.VarSpec `json:"client_vses"`
}

type ExecSnap struct {
	ImplRef  implsem.SemRef   `json:"ref"`
	Title    string           `json:"title"`
	SubExecs []implsem.SemRef `json:"subs"`
}

type PollSpec struct {
	ExecID string `json:"exec_id"`
}
