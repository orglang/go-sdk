package procstep

import (
	"github.com/orglang/go-sdk/adt/implsem"
	"github.com/orglang/go-sdk/adt/procexp"
)

type StepSpec struct {
	ExecRef implsem.SemRef
	ProcES  procexp.ExpSpec `json:"proc_es"`
}

type stepKind string

const (
	MsgStep = stepKind("msg")
	SvcStep = stepKind("svc")
)

type StepRec struct {
	ExecRef implsem.SemRef `json:"exec_ref"`
	Msg     *MsgRec        `json:"msg"`
	Svc     *SvcRec        `json:"svc"`
}

type MsgRec struct {
}

type SvcRec struct {
}
