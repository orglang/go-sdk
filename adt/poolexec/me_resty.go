package poolexec

import (
	"fmt"

	"github.com/go-resty/resty/v2"

	"github.com/orglang/go-sdk/adt/implsem"
	"github.com/orglang/go-sdk/adt/poolcomm"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Create(spec ExecSpec) (implsem.SemRef, error) {
	var dto implsem.SemRef
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetBody(&spec).
		Post("/pools/execs")
	if err != nil {
		return implsem.SemRef{}, err
	}
	if res.IsError() {
		return implsem.SemRef{}, fmt.Errorf("received: %v", string(res.Body()))
	}
	return dto, nil
}

func (sdk *RestySDK) Retrieve(ref implsem.SemRef) (ExecSnap, error) {
	var dto ExecSnap
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetPathParam("id", ref.ImplID).
		Get("/pools/{id}")
	if err != nil {
		return ExecSnap{}, err
	}
	if res.IsError() {
		return ExecSnap{}, fmt.Errorf("received: %v", string(res.Body()))
	}
	return dto, nil
}

func (sdk *RestySDK) RetreiveRefs() ([]implsem.SemRef, error) {
	refs := []implsem.SemRef{}
	return refs, nil
}

func (sdk *RestySDK) Take(spec poolcomm.CommSpec) error {
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

func (sdk *RestySDK) Spawn(spec poolcomm.CommSpec) (implsem.SemRef, error) {
	var dto implsem.SemRef
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetBody(&spec).
		Post("/pools/execs/spawns")
	if err != nil {
		return implsem.SemRef{}, err
	}
	if res.IsError() {
		return implsem.SemRef{}, fmt.Errorf("received: %v", string(res.Body()))
	}
	return dto, nil
}

func (sdk *RestySDK) Poll(spec PollSpec) (implsem.SemRef, error) {
	return implsem.SemRef{}, nil
}
