package pooldec

import (
	"github.com/orglang/go-sdk/adt/procbind"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type DecSpec struct {
	PoolNS               string
	PoolSN               string
	InsiderProvisionBC   procbind.BindSpec
	InsiderReceptionBC   procbind.BindSpec
	OutsiderProvisionBCs []procbind.BindSpec
	OutsiderReceptionBCs []procbind.BindSpec
}

type DecRef = uniqref.Msg
