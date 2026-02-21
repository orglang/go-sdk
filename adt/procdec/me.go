package procdec

import (
	"github.com/orglang/go-sdk/adt/descsem"
	"github.com/orglang/go-sdk/adt/descvar"
)

type DecSpec struct {
	DescQN     string            `json:"desc_qn"`
	ProviderVS descvar.VarSpec   `json:"provider_vs"`
	ClientVSes []descvar.VarSpec `json:"client_vses"`
}

type DecSnap struct {
	DescRef    descsem.SemRef   `json:"ref"`
	ProviderVR descvar.VarRec   `json:"provider_vr"`
	ClientVRs  []descvar.VarRec `json:"client_vrs"`
}
