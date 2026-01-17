package uniqsym

import (
	"github.com/orglang/go-sdk/adt/symbol"
)

const (
	regex = symbol.Chars + `+([.]` + symbol.Chars + `+)*`
)
