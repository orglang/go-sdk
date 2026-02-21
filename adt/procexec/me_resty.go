package procexec

import (
	"github.com/go-resty/resty/v2"

	"github.com/orglang/go-sdk/adt/implsem"
	"github.com/orglang/go-sdk/adt/procstep"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Take(spec procstep.StepSpec) error {
	var res implsem.SemRef
	_, err := sdk.Client.R().
		SetResult(&res).
		SetBody(&spec).
		SetPathParam("id", spec.ExecRef.ImplID).
		Post("/procs/{id}/steps")
	if err != nil {
		return err
	}
	return nil
}

func (sdk *RestySDK) Retrieve(execRef implsem.SemRef) (ExecSnap, error) {
	var res ExecSnap
	_, err := sdk.Client.R().
		SetPathParam("id", execRef.ImplID).
		SetResult(&res).
		Get("/procs/{id}")
	if err != nil {
		return ExecSnap{}, err
	}
	return res, nil
}
