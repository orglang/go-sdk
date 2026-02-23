package xactdef

import (
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/orglang/go-sdk/adt/descsem"
)

// Client-side secondary adapter
type RestySDK struct {
	Client *resty.Client
}

func (sdk *RestySDK) Incept(xactQN string) (descsem.SemRef, error) {
	return descsem.SemRef{}, nil
}

func (sdk *RestySDK) Create(spec DefSpec) (descsem.SemRef, error) {
	var dto descsem.SemRef
	res, err := sdk.Client.R().
		SetResult(&dto).
		SetBody(&spec).
		Post("/xacts/defs")
	if err != nil {
		return descsem.SemRef{}, err
	}
	if res.IsError() {
		return descsem.SemRef{}, fmt.Errorf("received: %v", string(res.Body()))
	}
	return dto, nil
}

func (sdk *RestySDK) Modify(snap DefSnap) (DefSnap, error) {
	return DefSnap{}, nil
}

func (sdk *RestySDK) Retrieve(defID string) (DefSnap, error) {
	return DefSnap{}, nil
}

func (sdk *RestySDK) RetreiveRefs() ([]descsem.SemRef, error) {
	return []descsem.SemRef{}, nil
}
