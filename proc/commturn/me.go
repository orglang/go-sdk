package commturn

import (
	"github.com/orglang/go-sdk/adt/commsem"
)

type turnKind string

const (
	PubKind turnKind = "pub"
	SubKind turnKind = "sub"
)

type TurnRec struct {
	CommRef commsem.SemRef `json:"comm_ref"`
	Pub     *PubRec        `json:"pub"`
	Sub     *SubRec        `json:"sub"`
}

type PubRec struct {
}

type SubRec struct {
}
