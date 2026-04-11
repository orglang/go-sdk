package procstep

import (
	"github.com/orglang/go-sdk/adt/procexp"
	"github.com/orglang/go-sdk/adt/semterm"
)

type StepSpec struct {
	ImplRef semterm.TermRef `json:"impl_ref"`
	ProcExp procexp.ExpSpec `json:"proc_exp"`
}

type stepKind string

const (
	Pub stepKind = "pub"
	Sub stepKind = "sub"
)

type StepRec struct {
	ImplRef semterm.TermRef `json:"impl_ref"`
	Pub     *PubRec         `json:"pub"`
	Sub     *SubRec         `json:"sub"`
}

type PubRec struct {
}

type SubRec struct {
}
