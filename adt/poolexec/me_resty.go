package poolexec

import (
	"github.com/go-resty/resty/v2"

	"github.com/orglang/go-sdk/adt/procexec"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Create(spec ExecSpecME) (ExecRefME, error) {
	var res ExecRefME
	_, err := sdk.Client.R().
		SetResult(&res).
		SetBody(&spec).
		Post("/pools")
	if err != nil {
		return ExecRefME{}, err
	}
	return res, nil
}

func (sdk *RestySDK) Retrieve(execID string) (ExecSnapME, error) {
	var res ExecSnapME
	_, err := sdk.Client.R().
		SetResult(&res).
		SetPathParam("id", execID).
		Get("/pools/{id}")
	if err != nil {
		return ExecSnapME{}, err
	}
	return res, nil
}

func (sdk *RestySDK) RetreiveRefs() ([]ExecRefME, error) {
	refs := []ExecRefME{}
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

func (sdk *RestySDK) Poll(spec PollSpecME) (procexec.ExecRef, error) {
	return procexec.ExecRef{}, nil
}
