package procdec

import (
	"github.com/orglang/go-sdk/adt/procbind"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type DecSpec struct {
	ProcQN     string              `json:"proc_qn"`
	ProviderBS procbind.BindSpec   `json:"provider_bs"`
	ClientBSes []procbind.BindSpec `json:"client_bses"`
}

type DecRef = uniqref.Msg

type DecSnap struct {
	DecRef     DecRef              `json:"ref"`
	ProviderBS procbind.BindSpec   `json:"provider_bs"`
	ClientBSes []procbind.BindSpec `json:"client_bses"`
}
