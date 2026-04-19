package termexp

import (
	"fmt"
)

type expKind string

const (
	Close expKind = "close"
	Wait  expKind = "wait"
	Send  expKind = "send"
	Recv  expKind = "recv"
	Lab   expKind = "lab"
	Case  expKind = "case"
	Link  expKind = "link"
	Fwd   expKind = "fwd"
	Call  expKind = "call"
)

type ExpSpec struct {
	K     expKind    `json:"kind"`
	Close *CloseSpec `json:"close,omitempty"`
	Wait  *WaitSpec  `json:"wait,omitempty"`
	Send  *SendSpec  `json:"send,omitempty"`
	Recv  *RecvSpec  `json:"recv,omitempty"`
	Lab   *LabSpec   `json:"lab,omitempty"`
	Case  *CaseSpec  `json:"case,omitempty"`
	Call  *CallSpec  `json:"call,omitempty"`
	Fwd   *FwdSpec   `json:"fwd,omitempty"`
}

type CloseSpec struct {
	CommChnlPH string `json:"comm_ph"`
}

type WaitSpec struct {
	CommChnlPH string  `json:"comm_ph"`
	ContES     ExpSpec `json:"cont_es"`
}

type SendSpec struct {
	CommChnlPH string `json:"comm_ph"`
	ValChnlPH  string `json:"val_ph"`
}

type RecvSpec struct {
	CommChnlPH string  `json:"comm_ph"`
	BindChnlPH string  `json:"bind_ph"`
	ContES     ExpSpec `json:"cont_es"`
}

type LabSpec struct {
	CommChnlPH string `json:"comm_ph"`
	PatternQN  string `json:"pattern_qn"`
}

type CaseSpec struct {
	CommChnlPH string       `json:"comm_ph"`
	ContBSes   []BranchSpec `json:"cont_bses"`
}

type BranchSpec struct {
	PatternQN string  `json:"pattern_qn"`
	ContES    ExpSpec `json:"cont_es"`
}

type CallSpec struct {
	BindChnlPH string   `json:"bind_ph"`
	ProcTermQN string   `json:"proc_qn"`
	ValChnlPHs []string `json:"val_phs"`
	ContES     ExpSpec  `json:"cont_es"`
}

type FwdSpec struct {
	CommChnlPH string `json:"comm_ph"`
	ContChnlPH string `json:"cont_ph"`
}

func ErrUnexpectedExpKind(k expKind) error {
	return fmt.Errorf("unexpected exp kind: %v", k)
}
