package pooldec

import (
	"github.com/orglang/go-sdk/adt/poolbind"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type DecSpec struct {
	PoolQN     string
	ProviderBS poolbind.BindSpec
	ClientBS   poolbind.BindSpec
}

type DecRef = uniqref.Msg
