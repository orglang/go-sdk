package poolexec

import (
	"github.com/go-resty/resty/v2"

	"github.com/orglang/go-sdk/adt/implsem"
	"github.com/orglang/go-sdk/adt/poolstep"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Create(spec ExecSpec) (implsem.SemRef, error) {
	var res implsem.SemRef
	_, err := sdk.Client.R().
		SetResult(&res).
		SetBody(&spec).
		Post("/pools/execs")
	if err != nil {
		return implsem.SemRef{}, err
	}
	return res, nil
}

func (sdk *RestySDK) Retrieve(ref implsem.SemRef) (ExecSnap, error) {
	var res ExecSnap
	_, err := sdk.Client.R().
		SetResult(&res).
		SetPathParam("id", ref.ImplID).
		Get("/pools/{id}")
	if err != nil {
		return ExecSnap{}, err
	}
	return res, nil
}

func (sdk *RestySDK) RetreiveRefs() ([]implsem.SemRef, error) {
	refs := []implsem.SemRef{}
	return refs, nil
}

func (sdk *RestySDK) Take(spec poolstep.StepSpec) error {
	_, err := sdk.Client.R().
		SetBody(&spec).
		SetPathParam("id", spec.ExecRef.ImplID).
		Post("/pools/execs/{id}/steps")
	if err != nil {
		return err
	}
	return nil
}

func (sdk *RestySDK) Spawn(spec poolstep.StepSpec) (implsem.SemRef, error) {
	var res implsem.SemRef
	_, err := sdk.Client.R().
		SetResult(&res).
		SetBody(&spec).
		SetPathParam("id", spec.ExecRef.ImplID).
		Post("/pools/{id}/spawns")
	if err != nil {
		return implsem.SemRef{}, err
	}
	return res, nil
}

func (sdk *RestySDK) Poll(spec PollSpec) (implsem.SemRef, error) {
	return implsem.SemRef{}, nil
}
