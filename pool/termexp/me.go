package termexp

import (
	"fmt"

	"github.com/orglang/go-sdk/adt/compsem"
	"github.com/orglang/go-sdk/adt/termsem"
)

type ExpSpec struct {
	K       expKind      `json:"kind"`
	Hire    *HireSpec    `json:"hire"`
	Fire    *FireSpec    `json:"fire"`
	Apply   *ApplySpec   `json:"apply"`
	Quit    *QuitSpec    `json:"quit"`
	Acquire *AcquireSpec `json:"acquire"`
	Release *ReleaseSpec `json:"release"`
	Accept  *AcceptSpec  `json:"accept"`
	Detach  *DetachSpec  `json:"detach"`
	Spawn   *SpawnSpec   `json:"spawn"`
}

type expKind string

const (
	Hire    expKind = "hire"
	Fire    expKind = "fire"
	Apply   expKind = "apply"
	Quit    expKind = "quit"
	Acquire expKind = "acquire"
	Release expKind = "release"
	Accept  expKind = "accept"
	Detach  expKind = "detach"
	Spawn   expKind = "spawn"
)

type HireSpec struct {
	CommChnlPH string  `json:"comm_ph"`
	ProcTermQN string  `json:"proc_qn"`
	ContExp    ExpSpec `json:"cont_exp"`
}

type FireSpec struct {
	CommChnlPH string
}

type ApplySpec struct {
	CommChnlPH string  `json:"comm_ph"`
	ProcTermQN string  `json:"proc_qn"`
	ContExp    ExpSpec `json:"cont_exp"`
}

type QuitSpec struct {
	CommChnlPH string `json:"comm_ph"`
	ProcDescQN string `json:"proc_qn"`
}

type AcquireSpec struct {
	CommChnlPH string  `json:"comm_ph"`
	ContExp    ExpSpec `json:"cont_exp"`
}

type ReleaseSpec struct {
	CommChnlPH string `json:"comm_ph"`
}

type AcceptSpec struct {
	CommChnlPH string  `json:"comm_ph"`
	ContExp    ExpSpec `json:"cont_exp"`
}

type DetachSpec struct {
	CommChnlPH string `json:"comm_ph"`
}

type SpawnSpec struct {
	// ссылка на описание порождаемого процесса
	ProcTermRef termsem.SemRef `json:"proc_dr"`
	// ссылки на воплощения потребляемых процессов
	ProcCompRefs []compsem.SemRef `json:"proc_irs"`
}

func ErrUnexpectedExpKind(k expKind) error {
	return fmt.Errorf("unexpected exp kind: %v", k)
}
