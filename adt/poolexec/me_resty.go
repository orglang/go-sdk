package poolexec

import (
	"fmt"
	"net/http"

	"github.com/go-resty/resty/v2"

	"github.com/orglang/go-sdk/adt/implsem"
	"github.com/orglang/go-sdk/adt/poolstep"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Create(spec ExecSpec) (implsem.SemRef, error) {
	var ref implsem.SemRef
	res, err := sdk.Client.R().
		SetResult(&ref).
		SetBody(&spec).
		Post("/pools/execs")
	if err != nil {
		return implsem.SemRef{}, err
	}
	if res.StatusCode() != http.StatusCreated {
		return implsem.SemRef{}, fmt.Errorf("creation failed")
	}
	return ref, nil
}

func (sdk *RestySDK) Retrieve(ref implsem.SemRef) (ExecSnap, error) {
	var snap ExecSnap
	res, err := sdk.Client.R().
		SetResult(&snap).
		SetPathParam("id", ref.ImplID).
		Get("/pools/{id}")
	if err != nil {
		return ExecSnap{}, err
	}
	if res.StatusCode() != http.StatusOK {
		return ExecSnap{}, fmt.Errorf("retrieval failed")
	}
	return snap, nil
}

func (sdk *RestySDK) RetreiveRefs() ([]implsem.SemRef, error) {
	refs := []implsem.SemRef{}
	return refs, nil
}

func (sdk *RestySDK) Take(spec poolstep.StepSpec) error {
	res, err := sdk.Client.R().
		SetBody(&spec).
		Post("/pools/execs/steps")
	if err != nil {
		return err
	}
	if res.StatusCode() != http.StatusNoContent {
		return fmt.Errorf("taking failed")
	}
	return nil
}

func (sdk *RestySDK) Spawn(spec poolstep.StepSpec) (implsem.SemRef, error) {
	var ref implsem.SemRef
	res, err := sdk.Client.R().
		SetResult(&ref).
		SetBody(&spec).
		Post("/pools/execs/spawns")
	if err != nil {
		return implsem.SemRef{}, err
	}
	if res.StatusCode() != http.StatusCreated {
		return implsem.SemRef{}, fmt.Errorf("taking failed")
	}
	return ref, nil
}

func (sdk *RestySDK) Poll(spec PollSpec) (implsem.SemRef, error) {
	return implsem.SemRef{}, nil
}
