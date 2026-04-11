package poolexec

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/orglang/go-sdk/adt/poolstep"
	"github.com/orglang/go-sdk/adt/semterm"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Create(spec ExecSpec) (semterm.TermRef, error) {
	var dto semterm.TermRef
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetBody(&spec).
		Post("/pools/execs")
	if err != nil {
		return semterm.TermRef{}, err
	}
	if res.IsError() {
		return semterm.TermRef{}, fmt.Errorf("received: %v", string(res.Body()))
	}
	return dto, nil
}

func (sdk *RestySDK) Retrieve(ref semterm.TermRef) (ExecSnap, error) {
	var dto ExecSnap
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetPathParam("id", ref.TermID).
		Get("/pools/{id}")
	if err != nil {
		return ExecSnap{}, err
	}
	if res.IsError() {
		return ExecSnap{}, fmt.Errorf("received: %v", string(res.Body()))
	}
	return dto, nil
}

func (sdk *RestySDK) RetreiveRefs() ([]semterm.TermRef, error) {
	refs := []semterm.TermRef{}
	return refs, nil
}

func (sdk *RestySDK) Take(spec poolstep.StepSpec) error {
	res, err := sdk.Client.R().
		SetBody(&spec).
		Post("/pools/execs/steps")
	if err != nil {
		return err
	}
	if res.IsError() {
		return fmt.Errorf("received: %v", string(res.Body()))
	}
	return nil
}

func (sdk *RestySDK) Spawn(spec poolstep.StepSpec) (semterm.TermRef, error) {
	var dto semterm.TermRef
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetBody(&spec).
		Post("/pools/execs/spawns")
	if err != nil {
		return semterm.TermRef{}, err
	}
	if res.IsError() {
		return semterm.TermRef{}, fmt.Errorf("received: %v", string(res.Body()))
	}
	return dto, nil
}
