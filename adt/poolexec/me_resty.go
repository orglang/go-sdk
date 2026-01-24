package poolexec

import (
	"github.com/go-resty/resty/v2"

	"github.com/orglang/go-sdk/adt/procexec"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Create(spec ExecSpec) (ExecRef, error) {
	var res ExecRef
	_, err := sdk.Client.R().
		SetResult(&res).
		SetBody(&spec).
		Post("/pools")
	if err != nil {
		return ExecRef{}, err
	}
	return res, nil
}

func (sdk *RestySDK) Retrieve(execID string) (ExecSnap, error) {
	var res ExecSnap
	_, err := sdk.Client.R().
		SetResult(&res).
		SetPathParam("id", execID).
		Get("/pools/{id}")
	if err != nil {
		return ExecSnap{}, err
	}
	return res, nil
}

func (sdk *RestySDK) RetreiveRefs() ([]ExecRef, error) {
	refs := []ExecRef{}
	return refs, nil
}

func (sdk *RestySDK) Spawn(spec procexec.ExecSpec) (procexec.ExecRef, error) {
	var res procexec.ExecRef
	_, err := sdk.Client.R().
		SetResult(&res).
		SetBody(&spec).
		SetPathParam("poolID", spec.PoolID).
		Post("/pools/{poolID}/procs")
	if err != nil {
		return procexec.ExecRef{}, err
	}
	return res, nil
}

func (sdk *RestySDK) Poll(spec PollSpec) (procexec.ExecRef, error) {
	return procexec.ExecRef{}, nil
}
