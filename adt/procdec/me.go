package procdec

import (
	"github.com/orglang/go-sdk/adt/descsem"
	"github.com/orglang/go-sdk/adt/implvar"
)

type DecSpec struct {
	ProcQN     string            `json:"proc_qn"`
	ProviderVS implvar.VarSpec   `json:"provider_vs"`
	ClientVSes []implvar.VarSpec `json:"client_vses"`
}

type DecSnap struct {
	DescRef    descsem.SemRef    `json:"ref"`
	ProviderVS implvar.VarSpec   `json:"provider_vs"`
	ClientVSes []implvar.VarSpec `json:"client_vses"`
}
