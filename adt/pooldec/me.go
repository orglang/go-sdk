package pooldec

import (
	"github.com/orglang/go-sdk/adt/procbind"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type DecSpec struct {
	PoolQN             string
	InsiderProvisionBS procbind.BindSpec
	InsiderReceptionBS procbind.BindSpec
}

type DecRef = uniqref.Msg
