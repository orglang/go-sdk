package poolexec

import (
	"github.com/go-resty/resty/v2"

	"github.com/orglang/go-sdk/adt/poolstep"
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

func (sdk *RestySDK) Retrieve(ref ExecRef) (ExecSnap, error) {
	var res ExecSnap
	_, err := sdk.Client.R().
		SetResult(&res).
		SetPathParam("id", ref.ID).
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

func (sdk *RestySDK) Take(spec poolstep.StepSpec) error {
	_, err := sdk.Client.R().
		SetBody(&spec).
		SetPathParam("id", spec.ExecRef.ID).
		Post("/pools/{id}/steps")
	if err != nil {
		return err
	}
	return nil
}

func (sdk *RestySDK) Spawn(spec poolstep.StepSpec) (procexec.ExecRef, error) {
	var res procexec.ExecRef
	_, err := sdk.Client.R().
		SetResult(&res).
		SetBody(&spec).
		SetPathParam("id", spec.ExecRef.ID).
		Post("/pools/{id}/spawns")
	if err != nil {
		return procexec.ExecRef{}, err
	}
	return res, nil
}

func (sdk *RestySDK) Poll(spec PollSpec) (procexec.ExecRef, error) {
	return procexec.ExecRef{}, nil
}
