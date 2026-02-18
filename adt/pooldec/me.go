package pooldec

import (
	"github.com/orglang/go-sdk/adt/poolbind"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type DecSpec struct {
	PoolQN     string              `json:"pool_qn"`
	ProviderBS poolbind.BindSpec   `json:"provider_bs"`
	ClientBSes []poolbind.BindSpec `json:"client_bses"`
}

type PoolRef = uniqref.Msg
