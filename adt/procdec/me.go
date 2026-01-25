package procdec

import (
	"github.com/orglang/go-sdk/adt/procbind"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type DecSpec struct {
	ProcQN     string              `json:"proc_qn"`
	ProviderBS procbind.BindSpec   `json:"provider_bs"`
	ClientBSs  []procbind.BindSpec `json:"client_bss"`
}

type DecRef = uniqref.Msg

type DecSnap struct {
	Ref  DecRef  `json:"ref"`
	Spec DecSpec `json:"spec"`
}
