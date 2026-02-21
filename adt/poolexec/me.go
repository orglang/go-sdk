package poolexec

import (
	"github.com/orglang/go-sdk/adt/implsem"
	"github.com/orglang/go-sdk/adt/implsubst"
)

type ExecSpec struct {
	DescQN     string                `json:"desc_qn"`
	ProviderSS implsubst.SubstSpec   `json:"provider_ss"`
	ClientSSes []implsubst.SubstSpec `json:"client_sses"`
}

type ExecSnap struct {
	ExecRef  implsem.SemRef   `json:"ref"`
	Title    string           `json:"title"`
	SubExecs []implsem.SemRef `json:"subs"`
}

type PollSpec struct {
	ExecID string `json:"exec_id"`
}
