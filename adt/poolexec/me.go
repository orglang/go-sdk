package poolexec

import (
	"github.com/orglang/go-sdk/adt/implsem"
)

type ExecSpec struct {
	PoolQN string `json:"pool_qn"`
}

type ExecSnap struct {
	ExecRef  implsem.SemRef   `json:"ref"`
	Title    string           `json:"title"`
	SubExecs []implsem.SemRef `json:"subs"`
}

type PollSpec struct {
	ExecID string `json:"exec_id"`
}
