package poolexec

import "github.com/orglang/go-sdk/adt/uniqref"

type ExecSpec struct {
	PoolQN string `json:"pool_qn"`
	SupID  string `json:"sup_id"`
}

type ExecRef = uniqref.Msg

type ExecSnap struct {
	ExecRef  ExecRef   `json:"ref"`
	Title    string    `json:"title"`
	SubExecs []ExecRef `json:"subs"`
}

type PollSpec struct {
	ExecID string `json:"exec_id"`
}
