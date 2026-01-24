package procdef

import "github.com/orglang/go-sdk/adt/uniqref"

type DefRec struct {
	Ref DefRef
}

type DefRef = uniqref.Msg
