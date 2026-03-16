package procexec

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/orglang/go-sdk/adt/implsem"
	"github.com/orglang/go-sdk/adt/procstep"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Take(spec procstep.StepSpec) error {
	var dto implsem.SemRef
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetBody(&spec).
		SetPathParam("id", spec.ImplRef.ImplID).
		Post("/procs/{id}/steps")
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("received: %v", string(res.Body()))
	}
	return nil
}

func (sdk *RestySDK) Retrieve(execRef implsem.SemRef) (ExecSnap, error) {
	var dto ExecSnap
	res, err := sdk.Client.R().
		SetPathParam("id", execRef.ImplID).
		SetResult(&dto).
		Get("/procs/{id}")
	if err != nil {
		return ExecSnap{}, err
	}
	if res.IsError() {
		return ExecSnap{}, fmt.Errorf("received: %v", string(res.Body()))
	}
	return dto, nil
}
