package pooldec

import (
	"github.com/orglang/go-sdk/adt/termctx"
)

type DecSpecME struct {
	PoolNS               string
	PoolSN               string
	InsiderProvisionBC   termctx.BindClaim
	InsiderReceptionBC   termctx.BindClaim
	OutsiderProvisionBCs []termctx.BindClaim
	OutsiderReceptionBCs []termctx.BindClaim
}

type DecRefME struct{}
