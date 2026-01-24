package procexec

import (
	"github.com/orglang/go-sdk/adt/uniqref"
)

type ExecRef = uniqref.Msg

type ExecSnap struct {
	ExecRef ExecRef
}
