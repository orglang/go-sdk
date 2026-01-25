package procstep

import (
	"github.com/orglang/go-sdk/adt/procexp"
	"github.com/orglang/go-sdk/adt/uniqref"
)

type StepSpec struct {
	ExecRef ExecRef
	ProcES  procexp.ExpSpec `json:"proc_es"`
}

type ExecRef = uniqref.Msg

type stepKind string

const (
	MsgStep = stepKind("msg")
	SvcStep = stepKind("svc")
)

type StepRec struct {
	ExecRef uniqref.Msg `json:"exec_ref"`
	Msg     *MsgRec     `json:"msg"`
	Svc     *SvcRec     `json:"svc"`
}

type MsgRec struct {
}

type SvcRec struct {
}
