package procdec

import (
	"github.com/orglang/go-sdk/adt/procbind"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type DecSpec struct {
	ProviderBS procbind.BindSpec   `json:"provider_bs"`
	ProcQN     string              `json:"proc_qn"`
	ClientBSs  []procbind.BindSpec `json:"client_bss"`
}

type DecRef = uniqref.Msg

type DecSnap struct {
	Ref  DecRef  `json:"ref"`
	Spec DecSpec `json:"spec"`
}
