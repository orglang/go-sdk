package procdec

import (
	"github.com/orglang/go-sdk/adt/procbind"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type DecSpec struct {
	ProcQN     string              `json:"proc_qn"`
	ClientBSs  []procbind.BindSpec `json:"client_bss"`
	ProviderBS procbind.BindSpec   `json:"provider_bs"`
}

type DecRef = uniqref.Msg

type DecSnap struct {
	DecRef     DecRef              `json:"ref"`
	ClientBSs  []procbind.BindSpec `json:"client_bss"`
	ProviderBS procbind.BindSpec   `json:"provider_bs"`
}
